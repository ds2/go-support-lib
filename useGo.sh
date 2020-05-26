#!/usr/bin/env bash

#export GOPATH=/Users/dstrauss/git/ds2-golang-lib::/Users/dstrauss/git/ds2-golang-lib/src
#export GOPATH=`pwd`:`pwd`/sysinfo
export GOBIN=`pwd`/shellbin
export PKGDIR=`pwd`/pkg

echo "Perform cleanup"
go clean
rm -r pkg 2>>/dev/null
rm -r bin 2>>/dev/null
rm -r vendor 2>>/dev/null
# echo "Formatting the code.."
# go fmt github.com/ds2/golang/sysinfo
echo "GETting the imports.."
go get -pkgdir "$PKGDIR" -v ./...
echo "Building final stuff"
go build -pkgdir "$PKGDIR" -v -x -buildmode=archive ./...
# go build -v -x -buildmode=archive github.com/ds2/golang/sysinfo
go test -run ''
echo "Installing it"
go install -pkgdir "$PKGDIR" -v ./...
echo "Run Doc server.. see http://localhost:6060 now "
godoc -http=:6060
echo "Done"
