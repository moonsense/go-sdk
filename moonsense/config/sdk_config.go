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

// SDKConfig is used to configure the SDK for accessing the Moonsense Cloud.
type SDKConfig struct {
	// API secret token generated from the Moonsense Cloud web console
	SecretToken string
	// Root API domain (defaults to moonsense.cloud)
	RootDomain string
	// Protocol to use when connecting to the API (defaults to https)
	Protocol string
	// Default Moonsense Cloud Data Plane region to connect to
	DefaultRegion string
}
