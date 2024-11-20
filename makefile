.PHONY: schema/generate
schema/generate:
	go run github.com/ogen-go/ogen/cmd/ogen --target internal/oas -package oas --clean openapi.yaml

