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

package models

import (
	"github.com/moonsense/go-sdk/moonsense/api"
	"github.com/moonsense/go-sdk/moonsense/config"

	dataPlaneProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/data-plane"
	dataPlaneSDKProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/data-plane-sdk"
)

// A PaginatedSessionList provides a wrapper around the list sessions response which allows for
// paginated loading of the sessions.
type PaginatedSessionList struct {
	// The sessions contained on this page
	Sessions []*dataPlaneSDKProto.Session

	sessionListResponse *dataPlaneProto.SessionListResponse
	currentPageConfig   config.ListSessionConfig
	dataPlaneClient     *api.DataPlaneClient
}

// Creates a new PaginatedSessionList.
func NewPaginatedSessionList(sessionListResponse *dataPlaneProto.SessionListResponse,
	listSessionConfig config.ListSessionConfig,
	dataPlaneClient *api.DataPlaneClient) PaginatedSessionList {

	return PaginatedSessionList{
		Sessions:            sessionListResponse.Sessions,
		sessionListResponse: sessionListResponse,
		currentPageConfig:   listSessionConfig,
		dataPlaneClient:     dataPlaneClient,
	}
}

// Returns true when there are more sessions to be fetched, otherwise false.
func (ps *PaginatedSessionList) HasMore() bool {
	if ps.sessionListResponse == nil ||
		ps.sessionListResponse.Pagination == nil {
		return false
	}

	return ps.sessionListResponse.Pagination.NextPage != 0
}

// Fetches the next page of sessions.
func (ps *PaginatedSessionList) NextPage() (*PaginatedSessionList, *api.ApiErrorResponse) {
	if !ps.HasMore() {
		empty := PaginatedSessionList{}
		return &empty, nil
	}

	// Calculate the next page config based on the incoming session list and such...
	lastSession := ps.sessionListResponse.Sessions[len(ps.sessionListResponse.Sessions)-1]

	nextPageConfig := ps.currentPageConfig
	nextPageConfig.Until = lastSession.CreatedAt.AsTime()

	sessionList, err := ps.dataPlaneClient.ListSessions(nextPageConfig)
	if err != nil {
		return nil, err
	}

	paginatedSessionList := NewPaginatedSessionList(sessionList, nextPageConfig, ps.dataPlaneClient)

	return &paginatedSessionList, nil
}
