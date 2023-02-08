package sdk

import (
	"fmt"
	"os"

	commonProto "github.com/moonsense/go-sdk/sdk/models/pb/v2/common"
)

const (
	defaultRootDomain      = "moonsense.cloud"
	defaultProtocol        = "https"
	defaultDataPlaneRegion = "us-central1.gcp"
)

type Config struct {
	// API secret token generated from the Moonsense Cloud web console
	SecretToken string
	// Root API domain (defaults to moonsense.cloud)
	RootDomain string
	// Protocol to use when connecting to the API (defaults to https)
	Protocol string
	// Moonsense Cloud Data Plane region to connect to
	DataPlaneRegion string
}

// This is internal cuz we want to force you through the... Constructor, NewClient
type clientImpl struct {
	Config Config
}

// The interface that we define what all this client is going to do, similar to
// the Python SDK
type Client interface {
	WhoAmI() *commonProto.TokenSelfResponse
}

// My not so real Constructor
func NewClient(c Config) Client {
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
		Config: c,
	}
}

func (client *clientImpl) WhoAmI() *commonProto.TokenSelfResponse {
	// Do some stuff later.
	fmt.Printf("I have a config %s\n", client.Config.SecretToken)

	return nil
}
