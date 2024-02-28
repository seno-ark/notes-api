go-dev:
	go run cmd/api/main.go

go-test:
	go test ./... -v

go-mock:
	mockery --all --dir internal --output internal/mocks
