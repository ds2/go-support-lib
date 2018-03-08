#!/bin/bash

#export GOPATH=/Users/dstrauss/git/ds2-golang-lib::/Users/dstrauss/git/ds2-golang-lib/src
export GOPATH=`pwd`
export GOBIN=`pwd`/shellbin
export PKGDIR=`pwd`/pkg

echo "Perform cleanup"
go clean
rm -r pkg 2>>/dev/null
rm -r bin 2>>/dev/null
# echo "Formatting the code.."
# go fmt github.com/ds2/golang/sysinfo
echo "GETting the imports.."
go get -pkgdir $PKGDIR -v github.com/ds2/golang/sysinfo
# echo "Building final stuff"
#go build -pkgdir $PKGDIR -v -x -buildmode=archive github.com/ds2/golang/sysinfo
# go build -v -x -buildmode=archive github.com/ds2/golang/sysinfo
echo "Installing it"
go install -pkgdir $PKGDIR -v github.com/ds2/golang/sysinfo
echo "Run Doc server.."
godoc -http=:6060
echo "Done"