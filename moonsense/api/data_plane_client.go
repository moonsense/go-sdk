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
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/moonsense/go-sdk/moonsense/config"
	"github.com/moonsense/go-sdk/moonsense/models/pb/v2/bundle"
	commonProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/common"
	dataPlaneProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/data-plane"
	dataPlaneSDKProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/data-plane-sdk"
)

const (
	dataPlaneVersion = "/v2"
)

type DataPlaneClient struct {
	apiClient *ApiClient
	config    config.SDKConfig
}

func NewDataPlaneClient(c config.SDKConfig) *DataPlaneClient {
	baseUrl := c.Protocol + "://" + c.DefaultRegion + ".data-api." + c.RootDomain

	api := ApiClient{
		BaseUrl:              baseUrl,
		HeaderBasedAccessKey: "Bearer " + c.SecretToken,
		HttpClient:           &http.Client{},
	}

	return &DataPlaneClient{
		apiClient: &api,
		config:    c,
	}
}

// Private helper methods

func (client *DataPlaneClient) baseUrl(region string) string {
	if region == "" {
		return client.config.Protocol + "://" + client.config.DefaultRegion + ".data-api." + client.config.RootDomain
	}

	return client.config.Protocol + "://" + region + ".data-api." + client.config.RootDomain
}

func (client *DataPlaneClient) findRegionId(sessionId string) (*string, *ApiErrorResponse) {
	session, err := client.DescribeSession(sessionId, true)
	if err != nil {
		return nil, err
	}

	return &session.RegionId, nil
}

func (client *DataPlaneClient) resetBaseUrl() {
	client.apiClient.BaseUrl = client.baseUrl("")
}

// Public methods

func (client *DataPlaneClient) ListJourneys(listJourneyConfig config.ListJourneyConfig) (*dataPlaneProto.JourneyListResponse, *ApiErrorResponse) {
	params := url.Values{}

	journeysPerPage := listJourneyConfig.JourneysPerPage
	if journeysPerPage < 0 {
		journeysPerPage = config.DefaultJourneysPerPage
	} else if journeysPerPage > config.MaxJourneysPerPage {
		journeysPerPage = config.MaxJourneysPerPage
	}
	params.Add("per_page", strconv.Itoa(journeysPerPage))

	for _, platform := range listJourneyConfig.Platforms {
		params.Add("filter[platforms][]", platform.String())
	}

	if !listJourneyConfig.Since.IsZero() {
		params.Add("filter[min_created_at]", listJourneyConfig.Since.Format(time.RFC3339))
	}

	if !listJourneyConfig.Until.IsZero() {
		params.Add("filter[max_created_at]", listJourneyConfig.Until.Format(time.RFC3339))
	}

	var response dataPlaneProto.JourneyListResponse
	err := client.apiClient.Get(
		NewDynamicPathWithQueryParams(dataPlaneVersion+"/journeys", map[string]string{}, params),
		&response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (client *DataPlaneClient) DescribeJourney(journeyId string) (*dataPlaneProto.JourneyDetailResponse, *ApiErrorResponse) {
	var response dataPlaneProto.JourneyDetailResponse

	err := client.apiClient.Get(
		NewDynamicPath(dataPlaneVersion+"/journeys/:journeyId", map[string]string{"journeyId": journeyId}),
		&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (client *DataPlaneClient) ListSessions(listSessionConfig config.ListSessionConfig) (*dataPlaneProto.SessionListResponse, *ApiErrorResponse) {
	params := url.Values{}

	sessionsPerPage := listSessionConfig.SessionsPerPage
	if sessionsPerPage < 0 {
		sessionsPerPage = config.DefaultSessionsPerPage
	} else if sessionsPerPage > config.MaxSessionsPerPage {
		sessionsPerPage = config.MaxSessionsPerPage
	}
	params.Add("per_page", strconv.Itoa(sessionsPerPage))

	if listSessionConfig.JourneyId != "" {
		params.Add("filter[journey_id]", listSessionConfig.JourneyId)
	}

	if !listSessionConfig.Since.IsZero() {
		params.Add("filter[min_created_at]", listSessionConfig.Since.Format(time.RFC3339))
	}

	if !listSessionConfig.Until.IsZero() {
		params.Add("filter[max_created_at]", listSessionConfig.Until.Format(time.RFC3339))
	}

	for _, label := range listSessionConfig.Labels {
		params.Add("filter[labels][]", label)
	}

	for _, platform := range listSessionConfig.Platforms {
		params.Add("filter[platforms][]", platform.String())
	}

	var response dataPlaneProto.SessionListResponse
	err := client.apiClient.Get(
		NewDynamicPathWithQueryParams(dataPlaneVersion+"/sessions", map[string]string{}, params),
		&response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (client *DataPlaneClient) DescribeSession(sessionId string, minimal bool) (*dataPlaneSDKProto.Session, *ApiErrorResponse) {
	var response dataPlaneSDKProto.Session

	var view = "full"
	if minimal {
		view = "minimal"
	}

	params := url.Values{}
	params.Add("view", view)

	err := client.apiClient.Get(
		NewDynamicPathWithQueryParams(dataPlaneVersion+"/sessions/:sessionId", map[string]string{"sessionId": sessionId}, params),
		&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (client *DataPlaneClient) ListSessionFeatures(sessionId string, region *string) (*dataPlaneProto.SessionFeaturesResponse, *ApiErrorResponse) {
	// If region is nil, we need to first lookup the session to get the region to use
	var regionId *string

	if region != nil {
		regionId = region
	} else {
		resolvedId, err := client.findRegionId(sessionId)
		if err != nil {
			return nil, err
		}
		regionId = resolvedId
	}

	// Modify the base url for this call
	defer client.resetBaseUrl()
	client.apiClient.BaseUrl = client.baseUrl(*regionId)

	var response dataPlaneProto.SessionFeaturesResponse
	err := client.apiClient.Get(
		NewDynamicPath(dataPlaneVersion+"/sessions/:sessionId/features", map[string]string{"sessionId": sessionId}),
		&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (client *DataPlaneClient) ListSessionSignals(sessionId string, region *string) (*dataPlaneProto.SignalsResponse, *ApiErrorResponse) {
	// If region is nil, we need to first lookup the session to get the region to use
	var regionId *string

	if region != nil {
		regionId = region
	} else {
		resolvedId, err := client.findRegionId(sessionId)
		if err != nil {
			return nil, err
		}
		regionId = resolvedId
	}

	// Modify the base url for this call
	defer client.resetBaseUrl()
	client.apiClient.BaseUrl = client.baseUrl(*regionId)

	var response dataPlaneProto.SignalsResponse
	err := client.apiClient.Get(
		NewDynamicPath(dataPlaneVersion+"/sessions/:sessionId/signals", map[string]string{"sessionId": sessionId}),
		&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (client *DataPlaneClient) ReadSession(sessionId string, region *string) ([]*bundle.SealedBundle, *ApiErrorResponse) {
	// If region is nil, we need to first lookup the session to get the region to use
	var regionId *string

	if region != nil {
		regionId = region
	} else {
		resolvedId, err := client.findRegionId(sessionId)
		if err != nil {
			return nil, err
		}
		regionId = resolvedId
	}

	// Modify the base url for this call
	defer client.resetBaseUrl()
	client.apiClient.BaseUrl = client.baseUrl(*regionId)

	var response []*bundle.SealedBundle
	err := client.apiClient.Get(
		NewDynamicPath(dataPlaneVersion+"/sessions/:sessionId/bundles", map[string]string{"sessionId": sessionId}),
		&response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (client *DataPlaneClient) UpdateSessionLabels(sessionId string, labels []string) *ApiErrorResponse {
	var response commonProto.Empty

	var sessionLabels []*dataPlaneSDKProto.SessionLabel
	for _, label := range labels {
		sessionLabels = append(sessionLabels, &dataPlaneSDKProto.SessionLabel{Name: label})
	}

	request := dataPlaneSDKProto.SessionLabelCreateRequest{
		Labels: sessionLabels,
	}

	err := client.apiClient.Post(
		NewDynamicPath(dataPlaneVersion+"/sessions/:sessionId/labels", map[string]string{"sessionId": sessionId}),
		&request,
		&response)

	return err
}

func (client *DataPlaneClient) WhoAmI() (*commonProto.TokenSelfResponse, *ApiErrorResponse) {
	var response commonProto.TokenSelfResponse

	err := client.apiClient.Get(
		NewStaticPath(dataPlaneVersion+"/tokens/self"),
		&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
