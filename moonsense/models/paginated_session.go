package models

import (
	"github.com/moonsense/go-sdk/moonsense/api"

	dataPlaneProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/data-plane"
	dataPlaneSDKProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/data-plane-sdk"
)

type PaginatedSession struct {
	Sessions []*dataPlaneSDKProto.Session

	sessionListResponse *dataPlaneProto.SessionListResponse
}

func NewPaginatedSession(sessionListResponse *dataPlaneProto.SessionListResponse) PaginatedSession {
	return PaginatedSession{
		Sessions:            sessionListResponse.Sessions,
		sessionListResponse: sessionListResponse,
	}
}

func (ps *PaginatedSession) NumberOfSessions() int {
	return len(ps.sessionListResponse.Sessions)
}

func (ps *PaginatedSession) CurrentPageNumber() int32 {
	return ps.sessionListResponse.Pagination.CurrentPage
}

func (ps *PaginatedSession) PreviousPageNumber() int32 {
	return ps.sessionListResponse.Pagination.PreviousPage
}

func (ps *PaginatedSession) NextPageNumber() int32 {
	return ps.sessionListResponse.Pagination.NextPage
}

func (ps *PaginatedSession) NumberPerPage() int32 {
	return ps.sessionListResponse.Pagination.PerPage
}

func (ps *PaginatedSession) NextPage() (*PaginatedSession, *api.ApiErrorResponse) {
	return nil, nil
}
