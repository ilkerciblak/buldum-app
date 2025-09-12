build:
	@go build -C api/ -o ../bin/buldum-app
run: build
	@./bin/buldum-app
test:
	@go test -v ./...
silent-Test:
	@go test ./...