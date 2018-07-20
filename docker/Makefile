PROJECT_NAME := Pulumi Docker Helpers
NODE_MODULE_NAME := @pulumi/docker
include ../build/common.mk

VERSION := $(shell ../scripts/get-version)

export PATH := $(shell yarn bin 2>/dev/null):$(PATH)

TESTPARALLELISM := 20

build::
	tsc
	sed -e 's/\$${VERSION}/$(VERSION)/g' < package.json > bin/package.json
	cp ../README.md ../LICENSE bin/

lint::
	tslint -c ../tslint.json -p tsconfig.json

test_all::
	PATH=$(PULUMI_BIN):$(PATH) go test -v -cover -timeout 2h -parallel ${TESTPARALLELISM} .

test_fast::
	PATH=$(PULUMI_BIN):$(PATH) go test -short -v -cover -timeout 2h -parallel ${TESTPARALLELISM} .
