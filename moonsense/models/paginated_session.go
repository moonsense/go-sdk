package models

import (
	"github.com/moonsense/go-sdk/moonsense/api"
	"github.com/moonsense/go-sdk/moonsense/config"

	dataPlaneProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/data-plane"
	dataPlaneSDKProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/data-plane-sdk"
)

type PaginatedSession struct {
	Sessions []*dataPlaneSDKProto.Session

	sessionListResponse *dataPlaneProto.SessionListResponse
	nextPageConfig      config.ListSessionConfig
	dataPlaneClient     *api.DataPlaneClient
}

func NewPaginatedSession(sessionListResponse *dataPlaneProto.SessionListResponse,
	listSessionConfig config.ListSessionConfig,
	dataPlaneClient *api.DataPlaneClient) PaginatedSession {

	// Calculate the next page config based on the incoming session list and such...
	nextPageConfig := listSessionConfig

	if sessionListResponse.Pagination.NextPage != 0 {
		lastSession := sessionListResponse.Sessions[len(sessionListResponse.Sessions)-1]
		nextPageConfig.Until = lastSession.CreatedAt.AsTime()
	}

	return PaginatedSession{
		Sessions:            sessionListResponse.Sessions,
		sessionListResponse: sessionListResponse,
		nextPageConfig:      nextPageConfig,
		dataPlaneClient:     dataPlaneClient,
	}
}

func (ps *PaginatedSession) NumberOfSessions() int {
	if ps.sessionListResponse == nil {
		return 0
	}

	return len(ps.sessionListResponse.Sessions)
}

func (ps *PaginatedSession) NumberPerPage() int {
	return ps.nextPageConfig.SessionsPerPage
}

func (ps *PaginatedSession) HasMoreSessions() bool {
	if ps.sessionListResponse == nil ||
		ps.sessionListResponse.Pagination == nil {
		return false
	}

	return ps.sessionListResponse.Pagination.NextPage != 0
}

func (ps *PaginatedSession) NextPage() (*PaginatedSession, *api.ApiErrorResponse) {
	if !ps.HasMoreSessions() {
		empty := PaginatedSession{}
		return &empty, nil
	}

	sessionList, err := ps.dataPlaneClient.ListSessions(ps.nextPageConfig)
	if err != nil {
		return nil, err
	}

	paginatedSession := NewPaginatedSession(sessionList, ps.nextPageConfig, ps.dataPlaneClient)

	return &paginatedSession, nil
}
