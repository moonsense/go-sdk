package sdk

import (
	"net/http"
	"net/url"

	cfg "github.com/moonsense/go-sdk/sdk/config"
	commonProto "github.com/moonsense/go-sdk/sdk/models/pb/v2/common"
	dataPlaneSDKProto "github.com/moonsense/go-sdk/sdk/models/pb/v2/data-plane-sdk"
)

type DataPlaneClient struct {
	apiClient *ApiClient
}

func NewDataPlaneClient(c cfg.Config) *DataPlaneClient {
	baseUrl := c.Protocol + "://" + c.DataPlaneRegion + ".data-api." + c.RootDomain

	api := ApiClient{
		BaseUrl:              baseUrl,
		HeaderBasedAccessKey: "Bearer " + c.SecretToken,
		HttpClient:           &http.Client{},
	}

	return &DataPlaneClient{
		apiClient: &api,
	}
}

func (client *DataPlaneClient) WhoAmI() (*commonProto.TokenSelfResponse, *ApiErrorResponse) {
	var response commonProto.TokenSelfResponse

	err := client.apiClient.ProcessRequest(
		"GET",
		NewStaticPath("/v2/tokens/self"),
		nil, &response)
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

	err := client.apiClient.ProcessRequest(
		"GET",
		NewDynamicPathWithQueryParams("/v2/sessions/:sessionId", map[string]string{"sessionId": sessionId}, params),
		nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
