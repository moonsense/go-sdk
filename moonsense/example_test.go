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

package moonsense_test

import (
	"fmt"
	"time"

	"github.com/moonsense/go-sdk/moonsense"
	"github.com/moonsense/go-sdk/moonsense/config"

	commonProto "github.com/moonsense/go-sdk/moonsense/models/pb/v2/common"
)

func Example_listRegions() {
	// Create an SDKConfig with your secret token
	sdkConfig := config.SDKConfig{
		SecretToken: "<YOUR SECRET_TOKEN>",
	}

	// Create a Moonsense Client with the SDKConfig
	client := moonsense.NewClient(sdkConfig)

	// Fetch the regions and print them
	regions, err := client.ListRegions()
	if err != nil {
		fmt.Println("Error fetching region list")
		fmt.Println(err)
		return
	}

	fmt.Println(regions)
}

func Example_listSessions() {
	// Create an SDKConfig with your secret token
	sdkConfig := config.SDKConfig{
		SecretToken: "<YOUR SECRET_TOKEN>",
	}

	// Create a Moonsense Client with the SDKConfig
	client := moonsense.NewClient(sdkConfig)

	// Configure the session list criteria.
	//
	// For this example, we want to retrieve 50 sessions per page that were created either with the
	// Android or iOS SDKs with a label of "suspect_session" and were created during the month of January, 2023.
	sessionListConfig := config.ListSessionConfig{
		SessionsPerPage: 50,
		Labels:          []string{"suspect_session"},
		Platforms:       []commonProto.DevicePlatform{commonProto.DevicePlatform_ANDROID, commonProto.DevicePlatform_iOS},
		Since:           time.Date(2023, time.January, 1, 0, 0, 0, 0, time.Local),
		Until:           time.Date(2023, time.January, 31, 23, 59, 59, 999, time.Local),
	}

	// Fetch the first page of sessions with the specified configuration.
	paginatedSessions, err := client.ListSessions(sessionListConfig)
	for {
		if err != nil {
			fmt.Println("Error fetching session list")
			fmt.Println(err)
			break
		}

		// Iterate over the returned sessions, printing out a subset of information
		for _, session := range paginatedSessions.Sessions {
			fmt.Printf("SessionId: %s, %s - %s\n", session.SessionId, session.Metadata.Platform.String(), session.CreatedAt.AsTime())
		}

		// After iterating across the first page of sessions, check to see if there are more available for this
		// criteria, and if so, fetch the next page of sessions.
		if paginatedSessions.HasMore() {
			paginatedSessions, err = paginatedSessions.NextPage()
		} else {
			break
		}
	}
}

func Example_describeSession() {
	// Create an SDKConfig with your secret token
	sdkConfig := config.SDKConfig{
		SecretToken: "<YOUR SECRET_TOKEN>",
	}

	// Create a Moonsense Client with the SDKConfig
	client := moonsense.NewClient(sdkConfig)

	// Fetch the session information for the specified sessionID
	session, err := client.DescribeSession("sessionID", false)
	if err != nil {
		fmt.Println("Error fetching session")
		fmt.Println(err)
		return
	}

	fmt.Println(session)
}

func Example_listSessionFeatures() {
	// Create an SDKConfig with your secret token
	sdkConfig := config.SDKConfig{
		SecretToken: "<YOUR SECRET_TOKEN>",
	}

	// Create a Moonsense Client with the SDKConfig
	client := moonsense.NewClient(sdkConfig)

	// Fetch the Features for the specified session
	featuresList, err := client.ListSessionFeatures("some session id", nil)
	if err != nil {
		fmt.Println("Error fetching session features")
		fmt.Println(err)
		return
	}

	fmt.Println(featuresList)
}

func Example_listSessionSignals() {
	// Create an SDKConfig with your secret token
	sdkConfig := config.SDKConfig{
		SecretToken: "<YOUR SECRET_TOKEN>",
	}

	// Create a Moonsense Client with the SDKConfig
	client := moonsense.NewClient(sdkConfig)

	// Fetch the Signals for the specified session
	signalsList, err := client.ListSessionSignals("some session id", nil)
	if err != nil {
		fmt.Println("Error fetching session signals")
		fmt.Println(err)
		return
	}

	fmt.Println(signalsList)
}

func Example_readSession() {
	// Create an SDKConfig with your secret token
	sdkConfig := config.SDKConfig{
		SecretToken: "<YOUR SECRET_TOKEN>",
	}

	// Create a Moonsense Client with the SDKConfig
	client := moonsense.NewClient(sdkConfig)

	// Read the sealed bundles for the specified session
	sealedBundles, err := client.ReadSession("some session id", nil)
	if err != nil {
		fmt.Println("Error fetching sealed bundles for session")
		fmt.Println(err)
	}

	fmt.Println(sealedBundles)
}

func Example_updateSessionLabels() {
	// Create an SDKConfig with your secret token
	sdkConfig := config.SDKConfig{
		SecretToken: "<YOUR SECRET_TOKEN>",
	}

	// Create a Moonsense Client with the SDKConfig
	client := moonsense.NewClient(sdkConfig)

	sessionId := "some session id"

	// Fetch the session information for the specified sessionID
	session, err := client.DescribeSession(sessionId, false)
	if err != nil {
		fmt.Println("Error fetching session")
		fmt.Println(err)
		return
	}

	// Examine the current labels
	if len(session.Labels) > 0 {
		fmt.Println(session.Labels)
	}

	// UpdateSessionLabels() replaces the current labels with the provided set of labels
	newLabels := []string{"data_validation_success", "pointer_feature_generator_validation"}
	err = client.UpdateSessionLabels(sessionId, newLabels)
	if err != nil {
		fmt.Println("Error updating session labels")
		fmt.Println(err)
		return
	}

	fmt.Println("Labels updated successfully")
}

func Example_whoAmI() {
	// Create an SDKConfig with your secret token
	sdkConfig := config.SDKConfig{
		SecretToken: "<YOUR SECRET_TOKEN>",
	}

	// Create a Moonsense Client with the SDKConfig
	client := moonsense.NewClient(sdkConfig)

	response, err := client.WhoAmI()
	if err != nil {
		fmt.Println("Error fetching token information")
		fmt.Println(err)
		return
	}

	fmt.Println(response)
}
