
APP_NAME := reblog
BUILD_DIR := .

SWAG := swag
AIR := air
GO := go
GOFUMPT := gofumpt
ABIGEN := abigen

.PHONY: dev all apidoc fmt

all:
	$(GO) build -v -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME) .

dev:
	$(AIR)

apidoc:
	$(SWAG) init --parseDependency --parseInternal --output ./docs --generalInfo server/server.go

abi:
	$(ABIGEN) --abi ./internal/copyright/abi.json --pkg copyright --type Copyright --out ./internal/copyright/generated.go

fmt:
	$(GOFUMPT) -l -w .
	$(SWAG) fmt
