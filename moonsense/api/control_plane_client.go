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
