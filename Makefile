build:
	@go build -o telcsifarmer 

run: build
	@./telcsifarmer

test:
	@go test -v ./...