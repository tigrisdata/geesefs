#!/usr/bin/env bash
# Copyright 2022-2023 Tigris Data, Inc.

set -e

VERSION=1.23.6
ARCH=$(dpkg --print-architecture)
FN="go${VERSION}.linux-${ARCH}.tar.gz"

case "$ARCH" in
"amd64")
  SHA256="9379441ea310de000f33a4dc767bd966e72ab2826270e038e78b2c53c2e7802d"
  ;;
"arm64")
  SHA256="561c780e8f4a8955d32bf72e46af0b5ee5e0debe1e4633df9a03781878219202"
  ;;
*)
  echo "No supported architecture."
  exit 1
  ;;
esac

wget "https://go.dev/dl/$FN"
echo "$SHA256  $FN" | shasum -a 256 -c

mkdir -p /usr/local
tar -C /usr/local -xzf "$FN"
rm "$FN"

export PATH=$PATH:/usr/local/go/bin

