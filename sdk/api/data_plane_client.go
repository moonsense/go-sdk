package sdk

import (
	"net/http"

	cfg "github.com/moonsense/go-sdk/sdk/config"
	commonProto "github.com/moonsense/go-sdk/sdk/models/pb/v2/common"
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
