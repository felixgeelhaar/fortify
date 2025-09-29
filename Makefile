.PHONY: test test-race test-cover bench fmt lint tidy clean

# Run all tests
test:
	go test -v ./...

# Run tests with race detection
test-race:
	go test -race -v ./...

# Run tests with coverage
test-cover:
	go test -cover -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Run benchmarks
bench:
	go test -bench=. -benchmem ./...

# Format code
fmt:
	gofmt -w -s .

# Run linter (requires golangci-lint)
lint:
	golangci-lint run

# Tidy dependencies
tidy:
	go mod tidy
	go mod verify

# Clean build artifacts
clean:
	rm -f coverage.out coverage.html
	go clean -testcache