package api

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/moonsense/go-sdk/moonsense/config"
	commonProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/common"
	dataPlaneProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/data-plane"
	dataPlaneSDKProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/data-plane-sdk"
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

func (client *DataPlaneClient) ListSessions(listSessionConfig config.ListSessionConfig) (*dataPlaneProto.SessionListResponse, *ApiErrorResponse) {
	params := url.Values{}

	page := 1
	if listSessionConfig.Page != 0 {
		page = listSessionConfig.Page
	}
	params.Add("page", strconv.Itoa(page))

	sessionsPerPage := 50
	if listSessionConfig.SessionsPerPage != 0 {
		sessionsPerPage = listSessionConfig.SessionsPerPage
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
		NewDynamicPathWithQueryParams("/v2/sessions", map[string]string{}, params),
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
		NewDynamicPathWithQueryParams("/v2/sessions/:sessionId", map[string]string{"sessionId": sessionId}, params),
		&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (client *DataPlaneClient) ListSessionFeatures(sessionId string, region *string) (*dataPlaneProto.FeatureListResponse, *ApiErrorResponse) {
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

	var response dataPlaneProto.FeatureListResponse
	err := client.apiClient.Get(
		NewDynamicPath("/v2/sessions/:sessionId/features", map[string]string{"sessionId": sessionId}),
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
		NewDynamicPath("/v2/sessions/:sessionId/signals", map[string]string{"sessionId": sessionId}),
		&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (client *DataPlaneClient) ReadSession(sessionId string, region *string) *ApiErrorResponse {
	// If region is nil, we need to first lookup the session to get the region to use
	var regionId *string

	if region != nil {
		regionId = region
	} else {
		resolvedId, err := client.findRegionId(sessionId)
		if err != nil {
			return err
		}
		regionId = resolvedId
	}

	// Modify the base url for this call
	defer client.resetBaseUrl()
	client.apiClient.BaseUrl = client.baseUrl(*regionId)

	err := client.apiClient.Get(
		NewDynamicPath("/v2/sessions/:sessionId/bundles", map[string]string{"sessionId": sessionId}),
		nil)

	if err != nil {
		fmt.Println("Badness happened!")
		fmt.Println(err)
		return err
	}

	return nil
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
		NewDynamicPath("/v2/sessions/:sessionId/labels", map[string]string{"sessionId": sessionId}),
		&request,
		&response)

	return err
}

func (client *DataPlaneClient) WhoAmI() (*commonProto.TokenSelfResponse, *ApiErrorResponse) {
	var response commonProto.TokenSelfResponse

	err := client.apiClient.Get(
		NewStaticPath("/v2/tokens/self"),
		&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}