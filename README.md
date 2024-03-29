# Moonsense SDK for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/moonsense/go-sdk/moonsense/.svg)](https://pkg.go.dev/github.com/moonsense/go-sdk/moonsense/) [![Apache V2 License](https://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/moonsense/go-sdk/blob/main/LICENSE)

Simple Go SDK for the Moonsense Cloud API.

## Installation

Install the module using `go get`: 

```shell
go get github.com/moonsense/go-sdk
```

## Getting Started

Start by getting an API secret key by navigating to App in Console and creating a token. You will need to save the generated secret key to a secure place.

https://console.moonsense.cloud/dashboard

We recommend exporting the API secret key as an environment variable:

    export MOONSENSE_SECRET_TOKEN=...

You can then very easily list sessions and access the granular data:

```go
package main

import (
    "fmt"

    "github.com/moonsense/go-sdk/moonsense"
    "github.com/moonsense/go-sdk/moonsense/config"
)

func main() {
    sdkConfig := config.SDKConfig{
        SecretToken: "<YOUR SECRET_TOKEN>",
    }

    client := moonsense.NewClient(sdkConfig)

    paginatedSession, err := client.ListSessions(config.ListSessionConfig{})
    for {
        if err != nil {
            fmt.Println("Error getting session list")
            fmt.Println(err)
            break
        }

        for _, session := range paginatedSession.Sessions {
            fmt.Printf("SessionId: %s, %s - %s\n", session.SessionId, session.Metadata.Platform.String(), session.CreatedAt.AsTime())
        }

        if paginatedSession.HasMore() {
            paginatedSession, err = paginatedSession.NextPage()
        } else {
            break
        }
    }

    // Fetch features for a specific session. Specifying the 
    // region the session is in will result in a faster lookup
    features, err := client.ListSessionFeatures("<session-id>", "asia-south1.gcp")
	if err != nil {
		fmt.Println("Error fetching session features")
		fmt.Println(err)
		return
	}

	fmt.Println(features)

    // Fetch signals for a specific session. Specifying the 
    // region the session is in will result in a faster lookup
    signals, err := client.ListSessionSignals("<session-id>", "asia-south1.gcp")
	if err != nil {
		fmt.Println("Error fetching session signals")
		fmt.Println(err)
		return
	}

	fmt.Println(signals)
}
```

### Add Journey Feedback

 Adding feedback to journeys provides a mechanism for tracking some special details against a given journey. For example, if a journey is determined to contain fraud, the journey can be flagged as fraudulent using Journey Feedback.

 ```go
	err := client.AddJourneyFeedback(journeyId, &journeysProto.JourneyFeedback{
		FraudFeedback: &journeysProto.FraudFeedback{
			IsFraud:     true,
			ReportedAt:  timestamppb.Now(),
			FraudReason: "It was fraud because...",
		},
	})
 ```

See [example_test.go](https://github.com/moonsense/go-sdk/blob/main/moonsense/example_test.go) for further examples.

## Terms Of Service

The Moonsense Go SDK is distributed under the [Moonsense Terms Of Service](https://www.moonsense.io/terms-of-service).

## Support

Feel free to raise an [Issue](https://github.com/moonsense/go-sdk/issues) around bugs, usage, concerns or feedback.
