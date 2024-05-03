
build:
	@go build -o ./bin/satur-api ./cmd/api

run: build
	@./bin/satur-api

clean: 
	@rm -rf ./bin/*

deps:
	@go mod tidy

test:
	@go test ./...
