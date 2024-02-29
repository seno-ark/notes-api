go-build:
	go build -o bin/api cmd/api/main.go

go-dev:
	go run cmd/api/main.go

go-test:
	go test -v -cover ./...

go-mock:
	mockery --all --dir internal --output internal/mocks
