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

package moonsense

import (
	"os"

	"github.com/moonsense/go-sdk/moonsense/api"
	"github.com/moonsense/go-sdk/moonsense/config"
	"github.com/moonsense/go-sdk/moonsense/models"
	"github.com/moonsense/go-sdk/moonsense/models/pb/v2/bundle"
	commonProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/common"
	controlPlaneProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/control-plane"
	dataPlaneProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/data-plane"
	dataPlaneSDKProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/data-plane-sdk"
	journeyFeedbackProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/journeys"
)

const (
	defaultRootDomain = "moonsense.cloud"
	defaultProtocol   = "https"
	defaultRegion     = "us-central1.gcp"
)

type clientImpl struct {
	Config             config.SDKConfig
	controlPlaneClient *api.ControlPlaneClient
	dataPlaneClient    *api.DataPlaneClient
}

// The Client interface defines the interface for interacting with the Moonsense Cloud.
type Client interface {
	// ListRegions retrieves the list of Data Plane regions in the Moonsense Cloud.
	//
	// These regions are used for data ingest and storage. Data is encrypted while at-rest
	// and in-transit. Granular data never leaves a region.
	//
	// See: https://api.moonsense.cloud/v2/regions
	ListRegions() ([]*controlPlaneProto.DataPlaneRegion, *api.ApiErrorResponse)

	// ListJourneys lists the journeys for the current project.
	ListJourneys(listJourneyConfig config.ListJourneyConfig) (*models.PaginatedJourneyList, *api.ApiErrorResponse)

	// DescribeJourney returns the details of a journey with the specified journeyId. The journey
	// response includes the list of sessions associated with the journey.
	DescribeJourney(journeyId string) (*dataPlaneProto.JourneyDetailResponse, *api.ApiErrorResponse)

	// GetJourneyFeedback returns the feedback associated with a journey with the specified journeyId.
	GetJourneyFeedback(journeyId string) (*journeyFeedbackProto.JourneyFeedback, *api.ApiErrorResponse)

	// AddJourneyFeedback sets the feedback associated with a journey with the specified journeyId.
	// Feedback is additive. If the feedback type already exists, it will be overwritten. If the
	// feedback type does not exist, it will be added.
	AddJourneyFeedback(journeyId string, feedback *journeyFeedbackProto.JourneyFeedback) *api.ApiErrorResponse

	// ListSessions lists the sessions for the current project
	ListSessions(listSessionConfig config.ListSessionConfig) (*models.PaginatedSessionList, *api.ApiErrorResponse)

	// DescribeSession returns the details of a session with the specified sessionId. If minimal is set to true
	// only total values are returned for counters.
	DescribeSession(sessionId string, minimal bool) (*dataPlaneSDKProto.Session, *api.ApiErrorResponse)

	// ListSessionFeatures returns the features for the specified sessionId. If region is not provided, the
	// appropriate region will be looked up by calling DescribeSession first.
	ListSessionFeatures(sessionId string, region *string) (*dataPlaneProto.SessionFeaturesResponse, *api.ApiErrorResponse)

	// ListSessionSignals returns the signals for the specified sessionId. If region is not provided, the
	// appropriate region will be looked up by calling DescribeSession first.
	ListSessionSignals(sessionId string, region *string) (*dataPlaneProto.SignalsResponse, *api.ApiErrorResponse)

	// ReadSession reads all data points for the specified Session that were collected so far. If region is not
	// provided, the appropriate region will be looked up by calling DescribeSession first.
	ReadSession(sessionId string, region *string) ([]*bundle.SealedBundle, *api.ApiErrorResponse)

	// UpdateSessionLabels replaces the existing session labels with the provided labels for the specified sessionId.
	UpdateSessionLabels(sessionId string, labels []string) *api.ApiErrorResponse

	// WhoAmI describes the authentication token used to connect to the API
	WhoAmI() (*commonProto.TokenSelfResponse, *api.ApiErrorResponse)
}

// NewClient returns a new Moonsense Client with the provided SDKConfig.
// If SDKConfig.SecretToken is not provided, the environment will be checked for the "MOONSENSE_SECRET_TOKEN" variable and will
// use that as the secret token. If neither are provided, the SDK will panic().
func NewClient(c config.SDKConfig) Client {
	if c.SecretToken == "" {
		secretToken := os.Getenv("MOONSENSE_SECRET_TOKEN")
		if secretToken != "" {
			c.SecretToken = secretToken
		} else {
			panic("A Secret Token must either be provided in config.SecretToken or set as an environment variable `MOONSENSE_SECRET_TOKEN`")
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

func (client *clientImpl) ListJourneys(listJourneyConfig config.ListJourneyConfig) (*models.PaginatedJourneyList, *api.ApiErrorResponse) {
	journeyList, err := client.dataPlaneClient.ListJourneys(listJourneyConfig)
	if err != nil {
		return nil, err
	}

	paginatedJourneyList := models.NewPaginatedJourneyList(journeyList, listJourneyConfig, client.dataPlaneClient)

	return &paginatedJourneyList, nil
}

func (client *clientImpl) DescribeJourney(journeyId string) (*dataPlaneProto.JourneyDetailResponse, *api.ApiErrorResponse) {
	return client.dataPlaneClient.DescribeJourney(journeyId)
}

func (client *clientImpl) GetJourneyFeedback(journeyId string) (*journeyFeedbackProto.JourneyFeedback, *api.ApiErrorResponse) {
	return client.dataPlaneClient.GetJourneyFeedback(journeyId)
}

func (client *clientImpl) AddJourneyFeedback(journeyId string, feedback *journeyFeedbackProto.JourneyFeedback) *api.ApiErrorResponse {
	return client.dataPlaneClient.AddJourneyFeedback(journeyId, feedback)
}

func (client *clientImpl) ListSessions(listSessionConfig config.ListSessionConfig) (*models.PaginatedSessionList, *api.ApiErrorResponse) {
	sessionList, err := client.dataPlaneClient.ListSessions(listSessionConfig)
	if err != nil {
		return nil, err
	}

	paginatedSessionList := models.NewPaginatedSessionList(sessionList, listSessionConfig, client.dataPlaneClient)

	return &paginatedSessionList, nil
}

func (client *clientImpl) DescribeSession(sessionId string, minimal bool) (*dataPlaneSDKProto.Session, *api.ApiErrorResponse) {
	return client.dataPlaneClient.DescribeSession(sessionId, minimal)
}

func (client *clientImpl) ListSessionFeatures(sessionId string, region *string) (*dataPlaneProto.SessionFeaturesResponse, *api.ApiErrorResponse) {
	return client.dataPlaneClient.ListSessionFeatures(sessionId, region)
}

func (client *clientImpl) ListSessionSignals(sessionId string, region *string) (*dataPlaneProto.SignalsResponse, *api.ApiErrorResponse) {
	return client.dataPlaneClient.ListSessionSignals(sessionId, region)
}

func (client *clientImpl) UpdateSessionLabels(sessionId string, labels []string) *api.ApiErrorResponse {
	return client.dataPlaneClient.UpdateSessionLabels(sessionId, labels)
}

func (client *clientImpl) ReadSession(sessionId string, region *string) ([]*bundle.SealedBundle, *api.ApiErrorResponse) {
	return client.dataPlaneClient.ReadSession(sessionId, region)
}

func (client *clientImpl) WhoAmI() (*commonProto.TokenSelfResponse, *api.ApiErrorResponse) {
	return client.dataPlaneClient.WhoAmI()
}
