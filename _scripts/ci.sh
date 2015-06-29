#!/bin/bash

set -e

THIS_SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "${THIS_SCRIPT_DIR}/.."

#
# Script for Continuous Integration
#

set -v

go get github.com/kisielk/errcheck
go install github.com/kisielk/errcheck

errcheck -asserts=true -blank=true ./...

go test -v ./...
#
# ==> DONE - OK
#