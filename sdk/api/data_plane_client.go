package sdk

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	cfg "github.com/moonsense/go-sdk/sdk/config"
	commonProto "github.com/moonsense/go-sdk/sdk/models/pb/v2/common"
	dataPlaneProto "github.com/moonsense/go-sdk/sdk/models/pb/v2/data-plane"
	dataPlaneSDKProto "github.com/moonsense/go-sdk/sdk/models/pb/v2/data-plane-sdk"
)

type DataPlaneClient struct {
	apiClient *ApiClient
	config    cfg.Config
}

func NewDataPlaneClient(c cfg.Config) *DataPlaneClient {
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

func (client *DataPlaneClient) DownloadSession(sessionId string) *ApiErrorResponse {
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

func (client *DataPlaneClient) ListSessions(labels *[]string,
	journeyId *string,
	platforms *[]commonProto.DevicePlatform,
	since *time.Time,
	until *time.Time) {

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

	var response *dataPlaneProto.FeatureListResponse
	err := client.apiClient.Get(
		NewDynamicPath("/v2/sessions/:sessionId/features", map[string]string{"sessionId": sessionId}),
		&response)
	if err != nil {
		return nil, err
	}

	return response, nil
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

	var response *dataPlaneProto.SignalsResponse
	err := client.apiClient.Get(
		NewDynamicPath("/v2/sessions/:sessionId/signals", map[string]string{"sessionId": sessionId}),
		&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
