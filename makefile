.PHONY: schema/generate
schema/generate:
	go run github.com/ogen-go/ogen/cmd/ogen@v1.8.1 --target internal/oas -package oas --clean openapi.yaml

