# DS/2 Go Support Lib

Some helpful functions.

## Protocol Buffers

Install Protobuffers version via:

    go get -buildmode=default github.com/golang/protobuf

To generate the stubs, run inside the respective directory:

    protoc --go_out=. common/*.proto
    protoc --go_out=. k8s/*.proto
    protoc --go_out=. sysinfo/*.proto

## Build as library

Use:

    go build ./...
    
Tests via

    go test ./...
    
## Module Handling

    export GO111MODULE=on

### Updating to latest patch/minor

    go list -u -m all

## Releasing

    go mod tidy
    go test ./... # or to fully test all packages -> go test all
    