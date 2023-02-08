package sdk

import (
	"os"

	api "github.com/moonsense/go-sdk/sdk/api"
	cfg "github.com/moonsense/go-sdk/sdk/config"
	commonProto "github.com/moonsense/go-sdk/sdk/models/pb/v2/common"
	controlPlaneProto "github.com/moonsense/go-sdk/sdk/models/pb/v2/control-plane"
)

const (
	defaultRootDomain      = "moonsense.cloud"
	defaultProtocol        = "https"
	defaultDataPlaneRegion = "us-central1.gcp"
)

// This is internal cuz we want to force you through the... Constructor, NewClient
type clientImpl struct {
	Config             cfg.Config
	controlPlaneClient *api.ControlPlaneClient
	dataPlaneClient    *api.DataPlaneClient
}

// The interface that we define what all this client is going to do, similar to
// the Python SDK
type Client interface {
	ListRegions() (*controlPlaneProto.DataRegionsListResponse, *api.ApiErrorResponse)
	WhoAmI() (*commonProto.TokenSelfResponse, *api.ApiErrorResponse)
}

// My not so real Constructor
func NewClient(c cfg.Config) Client {
	if c.SecretToken == "" {
		// Check the environment to see if it is set?
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
	if c.DataPlaneRegion == "" {
		c.DataPlaneRegion = defaultDataPlaneRegion
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
