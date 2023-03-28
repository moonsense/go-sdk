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
)

// A PaginatedJourneyList provides a wrapper around the list journeys response which allows for
// paginated loading of the journeys.
type PaginatedJourneyList struct {
	// The journeys contained on this page
	Journeys []*dataPlaneProto.Journey

	journeyListResponse *dataPlaneProto.JourneyListResponse
	currentPageConfig   config.ListJourneyConfig
	dataPlaneClient     *api.DataPlaneClient
}

// Creates a new PaginatedJourneyList.
func NewPaginatedJourneyList(journeyListResponse *dataPlaneProto.JourneyListResponse,
	listJourneyConfig config.ListJourneyConfig,
	dataPlaneClient *api.DataPlaneClient) PaginatedJourneyList {

	return PaginatedJourneyList{
		Journeys:            journeyListResponse.Journeys,
		journeyListResponse: journeyListResponse,
		currentPageConfig:   listJourneyConfig,
		dataPlaneClient:     dataPlaneClient,
	}
}

// Returns true when there are more journeys to be fetched, otherwise false.
func (ps *PaginatedJourneyList) HasMore() bool {
	if ps.journeyListResponse == nil ||
		ps.journeyListResponse.Pagination == nil {
		return false
	}

	return ps.journeyListResponse.Pagination.NextPage != 0
}

// Fetches the next page of journeys.
func (ps *PaginatedJourneyList) NextPage() (*PaginatedJourneyList, *api.ApiErrorResponse) {
	if !ps.HasMore() {
		empty := PaginatedJourneyList{}
		return &empty, nil
	}

	// Calculate the next page config based on the incoming journey list and such...
	lastJourney := ps.journeyListResponse.Journeys[len(ps.journeyListResponse.Journeys)-1]

	nextPageConfig := ps.currentPageConfig
	nextPageConfig.Until = lastJourney.CreatedAt.AsTime()

	journeyList, err := ps.dataPlaneClient.ListJourneys(nextPageConfig)
	if err != nil {
		return nil, err
	}

	paginatedJourneyList := NewPaginatedJourneyList(journeyList, nextPageConfig, ps.dataPlaneClient)

	return &paginatedJourneyList, nil
}
