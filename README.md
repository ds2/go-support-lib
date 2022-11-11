[![pipeline status](https://gitlab.com/ds_2/go-support-lib/badges/master/pipeline.svg)](https://gitlab.com/ds_2/go-support-lib/commits/master)
[![coverage report](https://gitlab.com/ds_2/go-support-lib/badges/master/coverage.svg)](https://gitlab.com/ds_2/go-support-lib/commits/master)


# Go Support Lib

Some helpful functions.

## Using it in your project

Run:

    go get gitlab.com/ds_2/go-support-lib

## Protocol Buffers

Install Protobuffers version via:

    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

To generate the stubs, run inside the respective directory:

    protoc --go_out=. common/*.proto
    protoc --go_out=. docker/*.proto
    protoc --go_out=. k8s/*.proto
    protoc --go_out=. sysinfo/*.proto

## Build as library

Run:

    go build -race ./...

To deploy it locally:

    go install ./...

### Update dependencies

    export GO111MODULE=on
    go mod tidy
    go clean -modcache
    go list -u -m all

Use:

    go build -race ./...
    
Tests via

    go test ./...

## Releasing

### Update License headers

Install the client (if necessary):

    go install github.com/apache/skywalking-eyes/cmd/license-eye@latest

Via:

    podman run -it --rm -v $(pwd):/github/workspace docker.io/apache/skywalking-eyes header check
    podman run -it --rm -v $(pwd):/github/workspace docker.io/apache/skywalking-eyes header fix

### Check release branch

    go mod tidy -v # cleans all unneeded deps
    go mod verify
    podman run -it --rm -v $(pwd):/github/workspace docker.io/apache/skywalking-eyes dependency check
    go test ./... # or to fully test all packages -> go test all
