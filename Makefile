GO := go
GOARCH := arm
GOOS := $(shell go env GOOS)
GOARM := 7
BIN := gosms-ng_$(GOOS)_$(GOARCH)$(GOARM)


VERSION := $(shell git describe --abbrev=0 --dirty=-SNAPSHOT --tags)
BUILD_COMMIT := $(shell git rev-parse --short HEAD)
SWAG_VERSION := $(shell swag -v 2>/dev/null)
GOBINDATA_VERSION := $(shell go-bindata -version 2>/dev/null)

LDFLAGS := "-X github.com/jiangytcn/gosms-ng/cmd.Version=$(VERSION) \
	-X github.com/jiangytcn/gosms-ng/cmd.Vcs=$(BUILD_COMMIT)"

check:
ifdef GOBINDATA_VERSION
else
	@echo go-bindata Not found
	@echo GO111MODULE=off go get -u github.com/go-bindata/go-bindata/...
	$(shell GO111MODULE=off go get -u github.com/go-bindata/go-bindata/...)
endif

.PHONY: vendor
export GO111MODULE=on
vendor:
	$(GO) mod download
	$(GO) mod vendor

.PHONY: ui
ui:
	cd ui; yarn build

.PHONY: assets
assets: ui
	go-bindata -fs -pkg dash -prefix='dash' -o dash/templates_bindatafs.go dash dash/static/css dash/static/js

.PHONY: dev
dev: assets
	$(GO) run main.go server

.PHONY: version
version:
	@echo $(VERSION)

.PHONY: generate
generate:
	$(GO) generate

.PHONY: build
build: assets generate
	env GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) $(GO) build -o $(BIN) -ldflags $(LDFLAGS)

