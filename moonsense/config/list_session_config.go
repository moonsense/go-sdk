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

package config

import (
	"time"

	commonProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/common"
)

const (
	// Default number of sessions to return when not specified
	DefaultSessionsPerPage = 25

	// Maximum number of sessions to return per page
	MaxSessionsPerPage = 100
)

// ListSessionConfig allows for setting the criteria to be used when requesting a
// list of sessions.
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
