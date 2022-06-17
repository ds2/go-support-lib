#!/usr/bin/env bash

apt-get update
apt-get install -y protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
protoc --go_out=. common/*.proto
protoc --go_out=. docker/*.proto
protoc --go_out=. k8s/*.proto
protoc --go_out=. sysinfo/*.proto
