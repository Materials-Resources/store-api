# ==================================================================================== #
# Helpers
# ==================================================================================== #

.PHONY: help
## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	git diff --exit-code

# ==================================================================================== #
# Code Generation
# ==================================================================================== #

.PHONY: gen/buf
## gen/buf: Generate code from buf.gen.yaml specifications
gen/buf:
	buf generate


.PHONY: gen/oapi
## gen/oapi: Generate code from openapi.yaml using ogen
gen/oapi:
	go run github.com/ogen-go/ogen/cmd/ogen@v1.8.1 --config config.ogen.yaml --target internal/oas -package oas --clean api/openapi.yaml


# ==================================================================================== #
# Quality Control
# ==================================================================================== #
.PHONY: lint
## lint: Run quality control checks
lint:
	go vet ./...
	go fmt ./...
	golangci-lint run

# ==================================================================================== #
# Application
# ==================================================================================== #

.PHONY: app/serve
## app/serve: Serve the application
app/serve:
	go run ./cmd/server



# ==================================================================================== #
# Build
# ==================================================================================== #

.PHONY: build
## build: Build the application
build:
	CGO_ENABLED=0 go build -trimpath -o ./bin/server_$(GOOS)_$(GOARCH)$(EXTENSION) ./cmd/server
.PHONY: build/linux_amd64
build/linux_amd64:
	GOOS=linux GOARCH=amd64 $(MAKE) build
.PHONY: build/linux_arm64
build/linux_arm64:
	GOOS=linux GOARCH=arm64 $(MAKE) build
.PHONY: build/darwin_amd64
build/darwin_amd64:
	GOOS=darwin GOARCH=amd64 $(MAKE) build
.PHONY: build/darwin_arm64
build/darwin_arm64:
	GOOS=darwin GOARCH=arm64 $(MAKE) build

