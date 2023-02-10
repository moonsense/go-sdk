package sdk

import (
	"os"
	"time"

	api "github.com/moonsense/go-sdk/sdk/api"
	cfg "github.com/moonsense/go-sdk/sdk/config"
	models "github.com/moonsense/go-sdk/sdk/models"
	commonProto "github.com/moonsense/go-sdk/sdk/models/pb/v2/common"
	controlPlaneProto "github.com/moonsense/go-sdk/sdk/models/pb/v2/control-plane"
	dataPlaneProto "github.com/moonsense/go-sdk/sdk/models/pb/v2/data-plane"
	dataPlaneSDKProto "github.com/moonsense/go-sdk/sdk/models/pb/v2/data-plane-sdk"
)

const (
	defaultRootDomain = "moonsense.cloud"
	defaultProtocol   = "https"
	defaultRegion     = "us-central1.gcp"
)

type clientImpl struct {
	Config             cfg.SDKConfig
	controlPlaneClient *api.ControlPlaneClient
	dataPlaneClient    *api.DataPlaneClient
}

type Client interface {
	// ListRegions retrieves the list of Data Plane regions in the Moonsense Cloud.
	//
	// These regions are used for data ingest and storage. Data is encrypted while at-rest
	// and in-transit. Granular data never leaves a region.
	//
	// See: https://api.moonsense.cloud/v2/regions
	ListRegions() ([]*controlPlaneProto.DataPlaneRegion, *api.ApiErrorResponse)

	// ListSessions lists the sessions for the current project
	ListSessions(listSessionConfig cfg.ListSessionConfig) (*models.PaginatedSession, *api.ApiErrorResponse)

	// DescribeSession returns the details of a session with the specified sessionId. If minimal is set to true
	// only total values are returned for counters.
	DescribeSession(sessionId string, minimal bool) (*dataPlaneSDKProto.Session, *api.ApiErrorResponse)

	// ListSessionFeatures returns the features for the specified sessionId. If region is not provided, the
	// appropriate region will be looked up by calling DescribeSession first.
	ListSessionFeatures(sessionId string, region *string) (*dataPlaneProto.FeatureListResponse, *api.ApiErrorResponse)

	// ListSessionSignals returns the signals for the specified sessionId. If region is not provided, the
	// appropriate region will be looked up by calling DescribeSession first.
	ListSessionSignals(sessionId string, region *string) (*dataPlaneProto.SignalsResponse, *api.ApiErrorResponse)

	// DownloadSession downloads and consolidates all data ingested so far for the session with the specified sessionId
	// into a single file with one bundle JSON per line.
	DownloadSession(sessionId string, outputFile string) *api.ApiErrorResponse

	// UpdateSessionLabels replaces the existing session labels with the provided labels for the specified sessionId.
	UpdateSessionLabels(sessionId string, labels []string) *api.ApiErrorResponse

	// WhoAmI describes the authentication token used to connect to the API
	WhoAmI() (*commonProto.TokenSelfResponse, *api.ApiErrorResponse)
}

func NewClient(c cfg.SDKConfig) Client {
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

func (client *clientImpl) ListRegions() ([]*controlPlaneProto.DataPlaneRegion, *api.ApiErrorResponse) {
	return client.controlPlaneClient.ListRegions()
}

func (client *clientImpl) ListSessions(listSessionConfig cfg.ListSessionConfig) (*models.PaginatedSession, *api.ApiErrorResponse) {
	sessionList, err := client.dataPlaneClient.ListSessions(listSessionConfig)
	if err != nil {
		return nil, err
	}

	paginatedSession := models.NewPaginatedSession(sessionList)

	return &paginatedSession, nil
}

func (client *clientImpl) DescribeSession(sessionId string, minimal bool) (*dataPlaneSDKProto.Session, *api.ApiErrorResponse) {
	return client.dataPlaneClient.DescribeSession(sessionId, minimal)
}

func (client *clientImpl) ListSessionFeatures(sessionId string, region *string) (*dataPlaneProto.FeatureListResponse, *api.ApiErrorResponse) {
	return client.dataPlaneClient.ListSessionFeatures(sessionId, region)
}

func (client *clientImpl) ListSessionSignals(sessionId string, region *string) (*dataPlaneProto.SignalsResponse, *api.ApiErrorResponse) {
	return client.dataPlaneClient.ListSessionSignals(sessionId, region)
}

func (client *clientImpl) UpdateSessionLabels(sessionId string, labels []string) *api.ApiErrorResponse {
	return client.dataPlaneClient.UpdateSessionLabels(sessionId, labels)
}

func (client *clientImpl) DownloadSession(sessionId string, outputFile string) *api.ApiErrorResponse {
	return client.dataPlaneClient.DownloadSession(sessionId)
}

func (client *clientImpl) WhoAmI() (*commonProto.TokenSelfResponse, *api.ApiErrorResponse) {
	time.Now()
	return client.dataPlaneClient.WhoAmI()
}
