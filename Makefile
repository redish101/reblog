
APP_NAME := reblog
BUILD_DIR := .

SWAG := swag
AIR := air
GO := go
GOFUMPT := gofumpt

.PHONY: dev all apidoc fmt

all:
	$(GO) build -v -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME) .

dev:
	$(AIR)

apidoc:
	$(SWAG) init --parseDependency --parseInternal --output ./docs

fmt:
	$(GOFUMPT) -l -w .
	$(SWAG) fmt
