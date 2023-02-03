package sdk

import (
	dataPlaneSDK "github.com/moonsense/go-sdk/sdk/models/pb/v2/data-plane-sdk"
)

func NewSession() *dataPlaneSDK.Session {
	session := dataPlaneSDK.Session{
		JourneyId: "IIsJourney",
	}

	return &session
}
