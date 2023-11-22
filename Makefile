run:
	@go run ./cmd/main.go

test:
	@GOFLAGS="-count=1" go test -v -cover -race ./...