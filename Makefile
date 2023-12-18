# Run application in development mode
dev:
	air
test:
	go test -v ./...
# Generate or update swagger docs
swag:
	swag init --dir ./cmd/server,./internal/handlers,./internal/models