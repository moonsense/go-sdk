package api

import (
	"net/http"

	"github.com/moonsense/go-sdk/moonsense/config"

	controlPlaneProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/control-plane"
)

type ControlPlaneClient struct {
	apiClient *ApiClient
}

func NewControlPlaneClient(c config.SDKConfig) *ControlPlaneClient {
	baseUrl := c.Protocol + "://api." + c.RootDomain

	api := ApiClient{
		BaseUrl:              baseUrl,
		HeaderBasedAccessKey: "Bearer " + c.SecretToken,
		HttpClient:           &http.Client{},
	}

	return &ControlPlaneClient{
		apiClient: &api,
	}
}

func (client *ControlPlaneClient) ListRegions() ([]*controlPlaneProto.DataPlaneRegion, *ApiErrorResponse) {
	var response controlPlaneProto.DataRegionsListResponse

	err := client.apiClient.Get(
		NewStaticPath("/v2/regions"),
		&response)
	if err != nil {
		return nil, err
	}

	return response.Regions, nil
}
