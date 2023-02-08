package sdk

import (
	"os"

	api "github.com/moonsense/go-sdk/sdk/api"
	cfg "github.com/moonsense/go-sdk/sdk/config"
	commonProto "github.com/moonsense/go-sdk/sdk/models/pb/v2/common"
	controlPlaneProto "github.com/moonsense/go-sdk/sdk/models/pb/v2/control-plane"
	dataPlaneSDKProto "github.com/moonsense/go-sdk/sdk/models/pb/v2/data-plane-sdk"
)

const (
	defaultRootDomain = "moonsense.cloud"
	defaultProtocol   = "https"
	defaultRegion     = "us-central1.gcp"
)

type clientImpl struct {
	Config             cfg.Config
	controlPlaneClient *api.ControlPlaneClient
	dataPlaneClient    *api.DataPlaneClient
}

type Client interface {
	ListRegions() (*controlPlaneProto.DataRegionsListResponse, *api.ApiErrorResponse)
	WhoAmI() (*commonProto.TokenSelfResponse, *api.ApiErrorResponse)
	DescribeSession(sessionId string, minimal bool) (*dataPlaneSDKProto.Session, *api.ApiErrorResponse)
	UpdateSessionLabels(sessionId string, labels []string) (*commonProto.Empty, *api.ApiErrorResponse)
}

func NewClient(c cfg.Config) Client {
	if c.SecretToken == "" {
		secretToken := os.Getenv("MOONSENSE_SECRET_TOKEN")
		if secretToken != "" {
			c.SecretToken = secretToken
		} else {
			panic("A Secret Token must either be provided in Config.SecretToken or set as an environment variable `MOONSENSE_SECRET_TOKEN`")
		}
	}
	if c.RootDomain == "" {
		c.RootDomain = defaultRootDomain
	}
	if c.Protocol == "" {
		c.Protocol = defaultProtocol
	}
	if c.DefaultRegion == "" {
		c.DefaultRegion = defaultRegion
	}

	return &clientImpl{
		Config:             c,
		controlPlaneClient: api.NewControlPlaneClient(c),
		dataPlaneClient:    api.NewDataPlaneClient(c),
	}
}

func (client *clientImpl) ListRegions() (*controlPlaneProto.DataRegionsListResponse, *api.ApiErrorResponse) {
	return client.controlPlaneClient.ListRegions()
}

func (client *clientImpl) WhoAmI() (*commonProto.TokenSelfResponse, *api.ApiErrorResponse) {
	return client.dataPlaneClient.WhoAmI()
}

func (client *clientImpl) DescribeSession(sessionId string, minimal bool) (*dataPlaneSDKProto.Session, *api.ApiErrorResponse) {
	return client.dataPlaneClient.DescribeSession(sessionId, minimal)
}

func (client *clientImpl) UpdateSessionLabels(sessionId string, labels []string) (*commonProto.Empty, *api.ApiErrorResponse) {
	return client.dataPlaneClient.UpdateSessionLabels(sessionId, labels)
}
