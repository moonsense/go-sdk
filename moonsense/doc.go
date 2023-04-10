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

// Package moonsense is the official Moonsense SDK for the Go programming language.
//
// # Getting started
//
// The best way to get started working with the SDK is to use `go get` to add the
// SDK to your Go dependencies explicitly.
//
//	go get github.com/moonsense/go-sdk
//
// Once you have installed the SDK, you will need an API secret key. Navigate to your App in the Moonsense Console and create a token.
// You will need to save the generated secret key to a secure place.
//
// https://console.moonsense.cloud/dashboard
//
// We recommend exporting the API secret key as an environment variable:
//
//	export MOONSENSE_SECRET_TOKEN=...
//
// You can then very easily list sessions and access the granular data:
//
//	package main
//
//	import (
//	    "fmt"
//
//	    "github.com/moonsense/go-sdk/moonsense"
//	    "github.com/moonsense/go-sdk/moonsense/config"
//	)
//
//	func main() {
//	    sdkConfig := config.SDKConfig{
//	        SecretToken: "<YOUR SECRET_TOKEN>",
//	    }
//
//	    client := moonsense.NewClient(sdkConfig)
//
//	    paginatedSession, err := client.ListSessions(config.ListSessionConfig{})
//	    for {
//	        if err != nil {
//	            fmt.Println("Error getting session list")
//	            fmt.Println(err)
//	            break
//	        }
//
//	        for _, session := range paginatedSession.Sessions {
//	            fmt.Printf("SessionId: %s, %s - %s\n", session.SessionId, session.Metadata.Platform.String(), session.CreatedAt.AsTime())
//	        }
//
//	        if paginatedSession.HasMore() {
//	            paginatedSession, err = paginatedSession.NextPage()
//	        } else {
//	            break
//	        }
//	    }
//	}
package moonsense
