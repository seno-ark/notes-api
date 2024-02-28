
go-dev:
	go run cmd/api/main.go

go-mock:
	mockery --all --dir internal --output internal/mocks
