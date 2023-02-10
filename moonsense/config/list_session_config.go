package config

import (
	"time"

	commonProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/common"
)

type ListSessionConfig struct {
	// The number of Sessions to return per page
	SessionsPerPage int

	// A list of Labels to match
	Labels []string

	// The journey id to match
	JourneyId string

	// The list of platforms to match. If none are specified, all platforms will be returned
	Platforms []commonProto.DevicePlatform

	// The start time to match
	Since time.Time

	// The end time to match
	Until time.Time
}
