package models

import (
	"github.com/moonsense/go-sdk/moonsense/api"
	"github.com/moonsense/go-sdk/moonsense/config"

	dataPlaneProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/data-plane"
	dataPlaneSDKProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/data-plane-sdk"
)

type PaginatedSessionList struct {
	Sessions []*dataPlaneSDKProto.Session

	sessionListResponse *dataPlaneProto.SessionListResponse
	currentPageConfig   config.ListSessionConfig
	dataPlaneClient     *api.DataPlaneClient
}

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

func (ps *PaginatedSessionList) HasMoreSessions() bool {
	if ps.sessionListResponse == nil ||
		ps.sessionListResponse.Pagination == nil {
		return false
	}

	return ps.sessionListResponse.Pagination.NextPage != 0
}

func (ps *PaginatedSessionList) NextPage() (*PaginatedSessionList, *api.ApiErrorResponse) {
	if !ps.HasMoreSessions() {
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

	paginatedSession := NewPaginatedSessionList(sessionList, nextPageConfig, ps.dataPlaneClient)

	return &paginatedSession, nil
}
