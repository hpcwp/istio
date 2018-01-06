## Copyright 2017 Istio Authors
##
## Licensed under the Apache License, Version 2.0 (the "License");
## you may not use this file except in compliance with the License.
## You may obtain a copy of the License at
##
##     http://www.apache.org/licenses/LICENSE-2.0
##
## Unless required by applicable law or agreed to in writing, software
## distributed under the License is distributed on an "AS IS" BASIS,
## WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
## See the License for the specific language governing permissions and
## limitations under the License.

#-----------------------------------------------------------------------------
# Global Variables
#-----------------------------------------------------------------------------
ISTIO_GO := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
SHELL := /bin/bash

# Current version, updated after a release.
VERSION ?= "0.5.0"

# Make sure GOPATH is set based on the executing Makefile and workspace. Will override
# GOPATH from the env.
export GOPATH= $(shell cd ../../..; pwd)

export CGO_ENABLE=0

# OUT is the directory where dist artifacts and temp files will be created.
OUT=${GOPATH}/out

GO ?= go

# Compile for linux/amd64 by default.
export GOOS ?= linux
export GOARCH ?= amd64

# Optional file including user-specific settings (HUB, TAG, etc)
-include .istiorc


# @todo allow user to run for a single $PKG only?
PACKAGES := $(shell $(GO) list ./...)
GO_EXCLUDE := /vendor/|.pb.go|.gen.go
GO_FILES := $(shell find . -name '*.go' | grep -v -E '$(GO_EXCLUDE)')

# Environment for tests, the directory containing istio and deps binaries.
# Typically same as GOPATH/bin, so tests work seemlessly with IDEs.
export ISTIO_BIN=${GOPATH}/bin

hub = ""
tag = ""

HUB?=istio
ifneq ($(strip $(HUB)),)
	hub =-hub ${HUB}
endif

# If tag not explicitly set in users' .istiorc or command line, default to the git sha.
TAG ?= $(shell git rev-parse --verify HEAD)
export TAG
export HUB
ifneq ($(strip $(TAG)),)
	tag =-tag ${TAG}
endif

#-----------------------------------------------------------------------------
# Output control
#-----------------------------------------------------------------------------
VERBOSE ?= 0
V ?= $(or $(VERBOSE),0)
Q = $(if $(filter 1,$V),,@)
H = $(shell printf "\033[34;1m=>\033[0m")

.DEFAULT_GOAL := build

checkvars:
	@if test -z "$(TAG)"; then echo "TAG missing"; exit 1; fi
	@if test -z "$(HUB)"; then echo "HUB missing"; exit 1; fi

setup: pilot/platform/kube/config

#-----------------------------------------------------------------------------
# Target: depend
#-----------------------------------------------------------------------------
.PHONY: depend depend.update
.PHONY: depend.status depend.ensure depend.graph

# Pull depdendencies, based on the checked in Gopkg.lock file.
# Developers must manually call dep.update if adding new deps or to pull recent
# changes.
depend: init

depend.ensure: init

# Target to update the Gopkg.lock with latest versions.
# Should be run when adding any new dependency and periodically.
depend.update: ${GOPATH}/bin/dep; $(info $(H) ensuring dependencies are up to date...)
	dep ensure
	dep ensure -update
	cp Gopkg.lock vendor/Gopkg.lock

${GOPATH}/bin/dep:
	go get -u github.com/golang/dep/cmd/dep

Gopkg.lock: Gopkg.toml | ${GOPATH}/bin/dep ; $(info $(H) generating) @
	$(Q) ${GOPATH}/bin/dep ensure -update

depend.status: Gopkg.lock
	$(Q) dep status > vendor/dep.txt
	$(Q) dep status -dot > vendor/dep.dot

# Requires 'graphviz' package. Run as user
depend.view: depend.status
	cat vendor/dep.dot | dot -T png > vendor/dep.png
	display vendor/dep.pkg
  
lint:
	SKIP_INIT=1 bin/linters.sh

# Target run by the pre-commit script, to automate formatting and lint
# If pre-commit script is not used, please run this manually.
pre-commit: fmt lint

# Downloads envoy, based on the SHA defined in the base pilot Dockerfile
# Will also check vendor, based on Gopkg.lock
init:
	@bin/init.sh

#-----------------------------------------------------------------------------
# Target: precommit
#-----------------------------------------------------------------------------
.PHONY: precommit format check
.PHONY: fmt format.gofmt format.goimports
.PHONY: check.vet check.lint

precommit: format check
format: format.goimports
fmt: format.gofmt format.goimports # backward compatible with ./bin/fmt.sh
check: check.vet check.lint

format.gofmt: ; $(info $(H) formatting files with go fmt...)
	$(Q) gofmt -s -w $(GO_FILES)

format.goimports: ; $(info $(H) formatting files with goimports...)
	$(Q) goimports -w -local istio.io $(GO_FILES)

# @todo fail on vet errors? Currently uses `true` to avoid aborting on failure
check.vet: ; $(info $(H) running go vet on packages...)
	$(Q) $(GO) vet $(PACKAGES) || true

# @todo fail on lint errors? Currently uses `true` to avoid aborting on failure
# @todo remove _test and mock_ from ignore list and fix the errors?
check.lint: ; $(info $(H) running golint on packages...)
	$(eval LINT_EXCLUDE := $(GO_EXCLUDE)|_test.go|mock_)
	$(Q) for p in $(PACKAGES); do \
		golint $$p | grep -v -E '$(LINT_EXCLUDE)' ; \
	done || true;

# @todo gometalinter targets?

build: setup go-build

#-----------------------------------------------------------------------------
# Target: go build
#-----------------------------------------------------------------------------

.PHONY: go-build

# gobuild script uses custom linker flag to set the variables.
# Params: OUT VERSION_PKG SRC

.PHONY: pilot
pilot: vendor
	bin/gobuild.sh ${GOPATH}/bin/pilot-discovery istio.io/istio/pilot/tools/version ./pilot/cmd/pilot-discovery

.PHONY: pilot-agent
pilot-agent: vendor
	bin/gobuild.sh ${GOPATH}/bin/pilot-agent istio.io/istio/pilot/tools/version ./pilot/cmd/pilot-agent

.PHONY: istioctl
istioctl: vendor
	bin/gobuild.sh ${GOPATH}/bin/istioctl istio.io/istio/pilot/tools/version ./pilot/cmd/istioctl

.PHONY: sidecar-initializer
sidecar-initializer: vendor
	bin/gobuild.sh ${GOPATH}/bin/sidecar-initializer istio.io/istio/pilot/tools/version ./pilot/cmd/sidecar-initializer

.PHONY: mixs
mixs: vendor
	bin/gobuild.sh ${GOPATH}/bin/mixs istio.io/istio/mixer/pkg/version ./mixer/cmd/mixs

.PHONY: mixc
mixc: vendor
	CGO_ENABLED=0 go build ${GOSTATIC} -o ${GOPATH}/bin/mixc istio.io/istio/mixer/cmd/mixc

.PHONY: node-agent
node-agent: vendor
	bin/gobuild.sh ${GOPATH}/bin/node_agent istio.io/istio/security/cmd/istio_ca/version ./security/cmd/node_agent

.PHONY: istio-ca
istio-ca: vendor
	bin/gobuild.sh ${GOPATH}/bin/istio_ca istio.io/istio/security/cmd/istio_ca/version ./security/cmd/istio_ca

go-build: pilot istioctl pilot-agent sidecar-initializer mixs mixc node-agent istio-ca

#-----------------------------------------------------------------------------
# Target: go test
#-----------------------------------------------------------------------------

.PHONY: go-test localTestEnv testApps

GOTEST_PARALLEL ?= '-test.parallel=4'
GOTEST_P ?= -p 1
GOSTATIC = -ldflags '-extldflags "-static"'

testApps:
	CGO_ENABLED=0 go build ${GOSTATIC} -o ${GOPATH}/bin/server istio.io/istio/pilot/test/server
	CGO_ENABLED=0 go build ${GOSTATIC} -o ${GOPATH}/bin/client istio.io/istio/pilot/test/client
	CGO_ENABLED=0 go build ${GOSTATIC} -o ${GOPATH}/bin/eurekamirror istio.io/istio/pilot/test/eurekamirror

localTestEnv: testApps
	bin/testEnvLocalK8S.sh ensure

# Temp. disable parallel test - flaky consul test.
# https://github.com/istio/istio/issues/2318
.PHONY: pilot-test
pilot-test: pilot-agent
	go test ${GOTEST_P} ${T} ./pilot/...

.PHONY: mixer-test
mixer-test: mixs
	# Some tests use relative path "testdata", must be run from mixer dir
	(cd mixer; go test ${T} ${GOTEST_PARALLEL} ./...)

.PHONY: broker-test
broker-test: vendor
	go test ${T} ./broker/...

.PHONY: security-test
security-test:
	go test ${T} ./security/...

# Run coverage tests
go-test: pilot-test mixer-test security-test broker-test

#-----------------------------------------------------------------------------
# Target: Code coverage ( go )
#-----------------------------------------------------------------------------

.PHONY: pilot-coverage
pilot-coverage:
	bin/parallel-codecov.sh pilot

.PHONY: mixer-coverage
mixer-coverage:
	bin/parallel-codecov.sh mixer

.PHONY: broker-coverage
broker-coverage:
	bin/parallel-codecov.sh broker

.PHONY: security-coverage
security-coverage:
	bin/parallel-codecov.sh security

# Run coverage tests
coverage: pilot-coverage mixer-coverage security-coverage broker-coverage


#-----------------------------------------------------------------------------
# Target: precommit
#-----------------------------------------------------------------------------
.PHONY: clean
.PHONY: clean.go

clean: clean.go

clean.go: ; $(info $(H) cleaning...)
	$(eval GO_CLEAN_FLAGS := -i -r)
	$(Q) $(GO) clean $(GO_CLEAN_FLAGS)
	$(MAKE) clean -C mixer
	$(MAKE) clean -C pilot
	$(MAKE) clean -C security

test: setup go-test

# Build all prod and debug images
docker:
	time $(ISTIO_GO)/security/bin/push-docker ${hub} ${tag} -build-only
	time $(ISTIO_GO)/mixer/bin/push-docker ${hub} ${tag} -build-only
	time $(ISTIO_GO)/pilot/bin/push-docker ${hub} ${tag} -build-only

# Build docker images for pilot
docker.pilot-debug: pilot pilot-agent sidecar-initializer
	cp ${GOPATH}/bin/{pilot-discovery,pilot-agent,sidecar-initializer} pilot/docker
	time (cd pilot/docker && docker build -t ${HUB}/proxy_debug:${TAG} -f Dockerfile.proxy_debug .)
	time (cd pilot/docker && docker build -t ${HUB}/proxy_init:${TAG} -f Dockerfile.proxy_init .)
	time (cd pilot/docker && docker build -t ${HUB}/sidecar_initializer:${TAG} -f Dockerfile.sidecar_initializer .)
	time (cd pilot/docker && docker build -t ${HUB}/pilot:${TAG} -f Dockerfile.pilot .)

# Build all docker debug images
docker.debug: docker.pilot-debug mixs mixc node-agent istio-ca
	cp ${GOPATH}/bin/mixs mixer/docker
	cp docker/ca-certificates.tgz mixer/docker
	time (cd mixer/docker && docker build -t ${HUB}/mixer_debug:${TAG} -f Dockerfile.debug .)
	cp ${GOPATH}/bin/{istio_ca,node_agent} security/docker
	cp docker/ca-certificates.tgz security/docker/
	time (cd security/docker && docker build -t ${HUB}/istio-ca:${TAG} -f Dockerfile.istio-ca .)

docker.pilot-test: testApps
	cp ${GOPATH}/bin/{client,server,eurekamirror} pilot/docker
	time (cd pilot/docker && docker build -t ${HUB}/app:${TAG} -f Dockerfile.app .)
	time (cd pilot/docker && docker build -t ${HUB}/eurekamirror:${TAG} -f Dockerfile.eurekamirror .)

# Build extra docker debug images (examples, etc)
docker.test: docker.pilot-test
	cp -a mixer/example/servicegraph/js/viz mixer/example/servicegraph/docker
	bin/gobuild.sh ${GOPATH}/bin/servicegraph istio.io/istio/mixer/pkg/version ./mixer/example/servicegraph/cmd/server
	cp ${GOPATH}/bin/servicegraph mixer/example/servicegraph/docker
	time (cd mixer/example/servicegraph/docker && docker build -t ${HUB}/servicegraph_debug:${TAG} -f Dockerfile.debug .)
	# TODO: generate or checkin test CA and keys
	security/bin/gen-keys.sh
	time (cd security/docker && docker build -t ${HUB}/istio-ca:${TAG} -f Dockerfile.istio-ca-test .)
	time (cd security/docker && docker build -t ${HUB}/istio-ca:${TAG} -f Dockerfile.node-agent-test .)


push: checkvars
	$(ISTIO_GO)/bin/push $(HUB) $(TAG)

artifacts: docker
	@echo 'To be added'

pilot/platform/kube/config:
	touch $@

kubelink:
	ln -fs ~/.kube/config pilot/platform/kube/

.PHONY: artifacts build checkvars clean docker test setup push kubelink

#-----------------------------------------------------------------------------
# Target: environment and tools
#-----------------------------------------------------------------------------
.PHONY: show.env show.goenv

show.env: ; $(info $(H) environment variables...)
	$(Q) printenv

show.goenv: ; $(info $(H) go environment...)
	$(Q) $(GO) version
	$(Q) $(GO) env

# show makefile variables. Usage: make show.<variable-name>
show.%: ; $(info $* $(H) $($*))
	$(Q) true

#-----------------------------------------------------------------------------
# Target: artifacts and distribution
#-----------------------------------------------------------------------------


${OUT}/dist/Gopkg.lock:
	mkdir -p ${OUT}/dist
	cp Gopkg.lock ${OUT}/dist/

# Binary/built artifacts of the distribution
dist-bin: ${OUT}/dist/Gopkg.lock

dist: dist-bin

include .circleci/Makefile

.PHONY: docker.sidecar.deb sidecar.deb ${OUT}/istio-sidecar.deb

# Make the deb image using the CI/CD image and docker.
docker.sidecar.deb:
	(cd ${TOP}; docker run --rm -u $(shell id -u) -it \
        -v ${GOPATH}:${GOPATH} \
        -w ${PWD} \
        -e USER=${USER} \
		--entrypoint /usr/bin/make ${CI_HUB}/ci:${CI_VERSION} \
		sidecar.deb )


# Create the 'sidecar' deb, including envoy and istio agents and configs.
# This target uses a locally installed 'fpm' - use 'docker.sidecar.deb' to use
# the builder image.
# TODO: consistent layout, possibly /opt/istio-VER/...
sidecar.deb: ${OUT}/istio-sidecar.deb

${OUT}/istio-sidecar.deb:
	mkdir -p ${OUT}
	fpm -s dir -t deb -n istio-sidecar -p ${OUT}/istio-sidecar.deb --version ${VERSION} --iteration 1 -C ${GOPATH} -f \
	   --url http://istio.io  \
	   --license Apache \
	   --vendor istio.io \
	   --maintainer istio@istio.io \
	   --after-install tools/deb/postinst.sh \
	   --config-files /var/lib/istio/envoy/sidecar.env \
	   --config-files /var/lib/istio/envoy/envoy.json \
	   --description "Istio" \
	   src/istio.io/istio/tools/deb/istio-start.sh=/usr/local/bin/istio-start.sh \
	   src/istio.io/istio/tools/deb/istio-iptables.sh=/usr/local/bin/istio-iptables/sh \
	   src/istio.io/istio/tools/deb/istio.service=/lib/systemd/system/istio.service \
	   src/istio.io/istio/security/tools/deb/istio-auth-node-agent.service=/lib/systemd/system/istio-auth-node-agent.service \
	   bin/envoy=/usr/local/bin/envoy \
	   bin/pilot-agent=/usr/local/bin/pilot-agent \
	   bin/node_agent=/usr/local/istio/bin/node_agent \
	   src/istio.io/istio/tools/deb/sidecar.env=/var/lib/istio/envoy/sidecar.env \
	   src/istio.io/istio/tools/deb/envoy.json=/var/lib/istio/envoy/envoy.json


#-----------------------------------------------------------------------------
# Target: e2e tests
#-----------------------------------------------------------------------------
include tests/istio.mk

