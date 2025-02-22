#!/bin/bash

#set -o xtrace
set -o errexit

T=
if [ $# == 1 ]; then
    T="-check.f $1"
fi

if [ "$NO_PROXY" == "" ]; then
  . "$(dirname "$0")/run-proxy.sh"
fi

if [ "$TIMEOUT" != "" ]; then
  TIMEOUT="-timeout $TIMEOUT"
fi

# run test in `go test` local mode so streaming output works
cd core
CGO_ENABLED=1 go test -race -v "$TIMEOUT" -check.vv $T
exit $?
