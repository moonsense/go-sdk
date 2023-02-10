VERSION 0.6

FROM golang:1.16-alpine

cleanup:
    LOCALLY
    RUN rm -rf moonsense/models/pb

generate:
    BUILD +cleanup
    ARG DEFINITIONS_BRANCH=main
    COPY github.com/moonsense/definitions/proto:$DEFINITIONS_BRANCH+generate-go-sdk/go-sdk /go-sdk
    SAVE ARTIFACT /go-sdk/github.com/moonsense/go-sdk/moonsense/models/pb AS LOCAL moonsense/models/pb
