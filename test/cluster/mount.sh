#!/bin/bash

p=`dirname $0`
FS_BIN="${FS_BIN:-"$p/../../geesefs"}"

_mount() {
  local MNT_DIR=$1
  shift
  echo "=== Mount $MNT_DIR"
  "$FS_BIN" \
  --endpoint="$ENDPOINT" \
  --enable-mtime \
  --cluster \
  "$@" \
  "$BUCKET_NAME" \
  "$MNT_DIR" &
}

_umount() {
  local MNT_DIR=$1
  echo "=== Unmount $MNT_DIR"
  umount "$MNT_DIR"
  sleep 1
  until [[ $(ps -ef | grep "geesefs" | grep "$MNT_DIR" | wc -l) == 0 ]]; do
    echo "=== Unmount $MNT_DIR... still doing"
    sleep 1
  done
}
