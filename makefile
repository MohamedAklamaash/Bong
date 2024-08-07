run:
	@go run cmd/main.go

gotest:
	@go test ./... -v 

build:
	@go build -o bin/main cmd/main.go
