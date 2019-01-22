# DS/2 Go Support Lib

Some helpful functions.

## Protocol Buffers

Install Protobuffers version via:

    go get -buildmode=default github.com/golang/protobuf

To generate the stubs, run inside the respective directory:

    protoc --go_out=. common/*.proto
    protoc --go_out=. k8s/*.proto
    protoc --go_out=. sysinfo/*.proto
