VERSION=$(shell git describe --tags --always)

run-test: s3proxy.jar build-debug
	./test/run-tests.sh

run-xfstests: s3proxy.jar xfstests build-debug
	./test/run-xfstests.sh

run-cluster-test: s3proxy.jar build-debug
	./test/cluster/test_random.sh

xfstests:
	git clone --depth=1 https://github.com/kdave/xfstests
	cd xfstests && patch -p1 -l < ../test/xfstests.diff

s3proxy.jar:
	wget https://github.com/gaul/s3proxy/releases/download/s3proxy-1.8.0/s3proxy -O s3proxy.jar

get-deps: s3proxy.jar
	go get -t ./...
	/bin/bash scripts/install_build_deps.sh
	/bin/bash scripts/install_test_deps.sh

build:
	go build -ldflags "github.com/tigrisdata/geesefs/core/cfg.Version=$(VERSION)"

build-debug:
	CGO_ENABLED=1 go build -race -ldflags "-X main.Version=`git rev-parse HEAD`"

install:
	go install -ldflags "github.com/tigrisdata/geesefs/core/cfg.Version=$(VERSION)"

# Setup local development environment.
setup: get-deps
	git config core.hooksPath ./.gitconfig/hooks

.PHONY: protoc
protoc:
	protoc --go_out=. --experimental_allow_proto3_optional --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative core/pb/*.proto

clean:
	rm -f geesefs
	rm -f core/mount_GoofysTest.*log
	findmnt -t fuse.geesefs -n -o TARGET|xargs -r umount
