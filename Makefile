.PHONY: run test build fmt vet

run:
	go run ./cmd/gateway -config configs/routes.json

test:
	go test -race -cover ./...

build:
	go build -o bin/aegis ./cmd/gateway

fmt:
	gofmt -w .

vet:
	go vet ./...
