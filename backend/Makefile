build:
	@go build -C cmd/ -o ../bin/buldum-app
run: build
	@./bin/buldum-app
test:
	@go test -v ./...
silent-Test:
	@go test ./...
test-service:
	@go test ${options} ./service/${service}/internal/${thing}
goose-up-service:
	@goose -dir ./service/${service}/internal/infrastructure/sql/migrations/. up
goose-down-service:
	@goose -dir ./service/${service}/internal/infrastructure/sql/migrations/. down

service-create-file:
	@touch ./service/${service}/internal/${place}/${f}