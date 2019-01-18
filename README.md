# DS/2 Go Support Lib

Some helpful functions.

## Protocol Buffers

To generate the stubs, run inside the respective directory:

    protoc --go_out=. k8s/k8s_types.proto
    protoc --go_out=. sysinfo/types.proto
