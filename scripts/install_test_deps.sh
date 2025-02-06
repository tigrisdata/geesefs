#!/usr/bin/env bash
# Copyright 2022-2023 Tigris Data, Inc.

set -ex

export GO111MODULE=on

curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$(go env GOPATH)/bin" latest

if [ "$(uname -s)" = "Darwin" ]; then
  if command -v brew > /dev/null 2>&1; then
    brew install shellcheck
  fi
else
  sudo apt-get install -y shellcheck \
    s3cmd \
    util-linux \
    fuse3
fi
