// Copyright 2023 Moonsense, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/moonsense/go-sdk/moonsense/models/pb/v2/bundle"
	commonProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/common"
)

type ApiClient struct {
	BaseUrl              string
	ServiceName          string
	UrlBasedAccessKey    string
	HeaderBasedAccessKey string
	HttpClient           *http.Client
}

// Common error reasons
const (
	InvalidRequestError = "invalid_request_error"
	APIError            = "api_error"
)

// A method for tracking an API error and it's parts
type ApiErrorResponse struct {
	Error         error                     // The raw error from the HTTP request
	HttpResponse  *http.Response            // The raw response from the HTTP request
	StatusCode    int                       // The status code, possibly adjusted by the API client
	ErrorResponse commonProto.ErrorResponse // A pre-built ErrorResponse
}

// Generic Path representation
type RequestPath interface {
	Encode() string
	RawPath() string
}

// Static Paths just store the path
type StaticPath struct {
	Path string
}

func NewStaticPath(path string) *StaticPath {
	return &StaticPath{Path: path}
}

func (s *StaticPath) Encode() string {
	return s.Path
}

func (s *StaticPath) RawPath() string {
	return s.Path
}

type DynamicPath struct {
	DynamicPath string
	PathParams  map[string]string
	QueryParams url.Values
}

func NewDynamicPath(dynamicPath string, params map[string]string) *DynamicPath {
	return &DynamicPath{
		DynamicPath: dynamicPath,
		PathParams:  params,
	}
}

func NewDynamicPathWithQueryParams(dynamicPath string, params map[string]string, queryParams url.Values) *DynamicPath {
	return &DynamicPath{
		DynamicPath: dynamicPath,
		PathParams:  params,
		QueryParams: queryParams,
	}
}

func (d *DynamicPath) Encode() string {
	var args []string
	for key, value := range d.PathParams {
		args = append(args, ":"+key)
		args = append(args, url.PathEscape(value))
	}

	replacer := strings.NewReplacer(args...)
	absolutePath := replacer.Replace(d.DynamicPath)

	if len(d.QueryParams) > 0 {
		absolutePath = absolutePath + "?" + d.QueryParams.Encode()
	}

	return absolutePath
}

func (s *DynamicPath) RawPath() string {
	return s.DynamicPath
}

func (err *ApiErrorResponse) String() string {
	if err.Error != nil {
		return err.Error.Error()
	}

	return err.ErrorResponse.Message
}

func (apiClient *ApiClient) send(method string, relativePath RequestPath,
	request interface{}, headers map[string]string) (*http.Response, error) {
	requestBodyBytes := []byte{}
	var requestBody *bytes.Buffer
	var err error

	// Convert the request body
	if request != nil {
		protoVal, isProto := request.(protoreflect.ProtoMessage)
		byteBuffer, isBytes := request.(*bytes.Buffer)

		if isProto {
			requestBodyBytes, err = proto.Marshal(protoVal)
		} else if isBytes {
			requestBody = byteBuffer
		} else {
			requestBodyBytes, err = json.Marshal(request)
		}

		if err != nil {
			return nil, err
		}
	}

	// If the requestBody buffer wasn't sent directly in the request
	// build it from the marshaled data
	if requestBody == nil {
		requestBody = bytes.NewBuffer(requestBodyBytes)
	}

	path := apiClient.BaseUrl + relativePath.Encode()

	// set the url based access key if needed
	if apiClient.UrlBasedAccessKey != "" {

		// append the access key
		if strings.Contains(path, "?") {
			path += "&" + apiClient.UrlBasedAccessKey
		} else {
			path += "?" + apiClient.UrlBasedAccessKey
		}
	}

	req, err := http.NewRequest(method, path, requestBody)
	if err != nil {
		return nil, err
	}

	// Add the requested headers if there are any
	for header, value := range headers {
		req.Header.Set(header, value)
	}

	// Set default Content-Type to protobuf if none was set already
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/x-protobuf")
	}

	// Set default Accept to protobuf if none was set already
	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/x-protobuf")
	}

	if apiClient.HeaderBasedAccessKey != "" {
		req.Header.Add("Authorization", apiClient.HeaderBasedAccessKey)
	}

	resp, err := apiClient.HttpClient.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (apiClient *ApiClient) SendAndDeserialize(
	method string,
	path RequestPath,
	request interface{}, response interface{}) (*http.Response, error) {

	return apiClient.SendAndDeserializeWithHeaders(method, path, request, response, nil)
}

func (apiClient *ApiClient) SendAndDeserializeWithHeaders(
	method string,
	path RequestPath,
	request interface{}, response interface{}, headers map[string]string) (*http.Response, error) {

	apiResponse, err := apiClient.send(method, path, request, headers)
	if err != nil {
		return nil, err
	}
	defer apiResponse.Body.Close()

	if apiResponse.StatusCode >= 200 && apiResponse.StatusCode < 300 {
		if response != nil {
			// TODO: Revist when the API method supports returning a byte array
			if apiResponse.Header.Get("Content-Type") == "application/octet-stream" {
				outsideBundles, isOK := response.(*[]*bundle.SealedBundle)

				if isOK {
					*outsideBundles = apiClient.extractSealedBundles(apiResponse.Body)
				}

				return apiResponse, nil
			}

			protoVal, isProto := response.(protoreflect.ProtoMessage)

			if isProto {
				err = apiClient.deserializeProto(apiResponse.Body, protoVal)
			} else {
				err = apiClient.deserializeStruct(apiResponse.Body, response)
			}

			if err != nil {
				return nil, err
			}
		}
	} else {
		responseBody, err := ioutil.ReadAll(apiResponse.Body)

		errorMessage := "an api error occurred"
		if err == nil {
			errorMessage = string(responseBody)
		}

		return apiResponse, errors.New(errorMessage)
	}

	return apiResponse, nil
}

func (apiClient *ApiClient) deserializeProto(body io.Reader, response proto.Message) error {
	responseBody, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(responseBody, response); err != nil {
		return err
	}

	return nil
}

func (apiClient *ApiClient) deserializeStruct(body io.Reader, response interface{}) error {
	responseBody, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(responseBody, response); err != nil {
		return err
	}

	return nil
}

func (apiClient *ApiClient) Get(path RequestPath, response interface{}) *ApiErrorResponse {
	return apiClient.ProcessRequest("GET", path, nil, response)
}

func (apiClient *ApiClient) Post(path RequestPath, requestBody interface{}, response interface{}) *ApiErrorResponse {
	return apiClient.ProcessRequest("POST", path, requestBody, response)
}

func (apiClient *ApiClient) PostForm(path RequestPath, requestBody map[string]interface{}, response interface{}) *ApiErrorResponse {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range requestBody {
		var fw io.Writer
		var err error
		var reader io.Reader
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add a file
		if x, ok := r.(*bytes.Buffer); ok {
			if fw, err = w.CreateFormFile(key, "pcap"); err != nil {
				return &ApiErrorResponse{
					Error: err,
					ErrorResponse: commonProto.ErrorResponse{
						Type:    APIError,
						Message: "Error creating file field for form post",
					},
				}
			}

			reader = x
		} else if s, ok := r.(string); ok {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return &ApiErrorResponse{
					Error: err,
					ErrorResponse: commonProto.ErrorResponse{
						Type:    APIError,
						Message: "Error creating string field for form post",
					},
				}
			}

			reader = strings.NewReader(s)
		} else {
			return &ApiErrorResponse{
				Error: fmt.Errorf("unsupported field type passed to form upload"),
				ErrorResponse: commonProto.ErrorResponse{
					Type:    APIError,
					Message: "unsupported field type passed to form upload",
				},
			}

		}
		if _, err = io.Copy(fw, reader); err != nil {
			return &ApiErrorResponse{
				Error: err,
				ErrorResponse: commonProto.ErrorResponse{
					Type:    APIError,
					Message: "error copying data to form field",
				},
			}
		}

	}
	// Close the multipart writer
	w.Close()

	headers := map[string]string{
		"Content-Type": w.FormDataContentType(),
	}

	return apiClient.ProcessRequestWithHeaders("POST", path, &b, response, headers)
}

func (apiClient *ApiClient) ProcessRequest(method string,
	path RequestPath,
	request interface{}, response interface{}) *ApiErrorResponse {

	return apiClient.ProcessRequestWithHeaders(method, path, request, response, nil)
}

func (apiClient *ApiClient) ProcessRequestWithHeaders(method string,
	path RequestPath,
	request interface{}, response interface{}, headers map[string]string) *ApiErrorResponse {

	httpResponse, requestErr := apiClient.SendAndDeserializeWithHeaders(method, path, request, response, headers)
	if httpResponse == nil && requestErr != nil {
		return &ApiErrorResponse{
			Error:      requestErr,
			StatusCode: http.StatusBadGateway,
			ErrorResponse: commonProto.ErrorResponse{
				Type:    APIError,
				Param:   fmt.Sprintf("%s request", apiClient.ServiceName),
				Message: "Connection Error.",
			},
		}
	}

	if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		errorMessage := ""

		// Propagate the proper error code that the control plane requested.
		if requestErr != nil {
			errorMessage = requestErr.Error()
		}

		errorType := APIError
		if httpResponse.StatusCode%100 == 4 {
			errorType = InvalidRequestError
		}

		// Try to parse as an ErrorResponse
		if strings.HasPrefix(errorMessage, "{") {

			var response commonProto.ErrorResponse
			if err := proto.Unmarshal([]byte(errorMessage), &response); err == nil {
				errorMessage = response.Message
			}
		}

		return &ApiErrorResponse{
			Error:        requestErr,
			HttpResponse: httpResponse,
			StatusCode:   httpResponse.StatusCode,
			ErrorResponse: commonProto.ErrorResponse{
				Type:    errorType,
				Param:   fmt.Sprintf("%s request", apiClient.ServiceName),
				Message: errorMessage,
			},
		}
	}

	return nil
}

func (apiClient *ApiClient) extractSealedBundles(gzipStream io.Reader) []*bundle.SealedBundle {
	var sealedBundles []*bundle.SealedBundle

	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		log.Fatal("extractSealedBundles: NewReader failed")
	}

	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("extractSealedBundles: Next() failed: %s", err.Error())
		}

		if header.Typeflag == tar.TypeReg {
			bufferedReader := bufio.NewReader(tarReader)

			for {
				bytes, err := bufferedReader.ReadBytes('\n')

				if err == io.EOF {
					break
				} else if err != nil {
					fmt.Println(err)
				} else {
					var response bundle.SealedBundle
					err := protojson.Unmarshal(bytes, &response)

					if err != nil {
						fmt.Println(err)
					} else {
						sealedBundles = append(sealedBundles, &response)
					}
				}
			}
		}
	}

	return sealedBundles
}
