# Go Support Lib

Some helpful functions.

## Protocol Buffers

Install Protobuffers version via:

    go get -buildmode=default github.com/golang/protobuf

To generate the stubs, run inside the respective directory:

    protoc --go_out=. common/*.proto
    protoc --go_out=. docker/*.proto
    protoc --go_out=. k8s/*.proto
    protoc --go_out=. sysinfo/*.proto

## Build as library

### Update dependencies

    export GO111MODULE=on
    go clean -modcache
    go list -u -m all

Use:

    go build -race ./...
    
Tests via

    go test ./...

## Releasing

    go mod tidy
    go mod verify
    go test ./... # or to fully test all packages -> go test all
