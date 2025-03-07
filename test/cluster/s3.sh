#!/bin/bash

echo "=== Use S3 ==="

unset ENDPOINT

TEST_ARTIFACTS="${TEST_ARTIFACTS:-executions/$(date +"%Y-%m-%d-%H-%M-%S")}"
mkdir -p "$TEST_ARTIFACTS"

_kill_s3proxy() {
    :
}

#  --signature-v2 \
_s3cmd() {
  s3cmd \
  --host-bucket="" \
  --host="$ENDPOINT" \
  "$@"
}