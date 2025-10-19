.PHONY: help start stop logs test clean

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

start: ## Start the distributed rate limiter example
	@echo "Starting Redis and application instances..."
	docker-compose up -d
	@echo ""
	@echo "Services started:"
	@echo "  - Redis:    http://localhost:6379"
	@echo "  - App 1:    http://localhost:8081"
	@echo "  - App 2:    http://localhost:8082"
	@echo "  - App 3:    http://localhost:8083"
	@echo ""
	@echo "Test distributed rate limiting:"
	@echo "  make test"
	@echo ""
	@echo "View logs:"
	@echo "  make logs"

stop: ## Stop all services
	@echo "Stopping services..."
	docker-compose down

restart: stop start ## Restart all services

logs: ## View logs from all services
	docker-compose logs -f

logs-app1: ## View logs from app instance 1
	docker-compose logs -f app1

logs-app2: ## View logs from app instance 2
	docker-compose logs -f app2

logs-app3: ## View logs from app instance 3
	docker-compose logs -f app3

logs-redis: ## View Redis logs
	docker-compose logs -f redis

status: ## Show status of all services
	docker-compose ps

test: ## Test distributed rate limiting across instances
	@echo "Testing distributed rate limiting..."
	@echo ""
	@echo "Making 15 requests distributed across 3 instances:"
	@echo "(Rate limit: 10 requests/minute, Burst: 15)"
	@echo ""
	@for i in 1 2 3 4 5; do \
		echo "Request $$i to app1:"; \
		curl -s -H "X-User-ID: test-user" http://localhost:8081/api/data | jq -r '.message // .error' || echo "Failed"; \
		echo ""; \
	done
	@for i in 6 7 8 9 10; do \
		echo "Request $$i to app2:"; \
		curl -s -H "X-User-ID: test-user" http://localhost:8082/api/data | jq -r '.message // .error' || echo "Failed"; \
		echo ""; \
	done
	@for i in 11 12 13 14 15; do \
		echo "Request $$i to app3:"; \
		curl -s -H "X-User-ID: test-user" http://localhost:8083/api/data | jq -r '.message // .error' || echo "Failed"; \
		echo ""; \
	done
	@echo ""
	@echo "Test complete! The 11th-15th requests should be rate limited."

test-single: ## Test rate limiting on a single instance
	@echo "Testing rate limiting on single instance..."
	@echo "(Rate limit: 10 requests/minute)"
	@echo ""
	@for i in $$(seq 1 15); do \
		echo -n "Request $$i: "; \
		curl -s -H "X-User-ID: test-user" http://localhost:8081/api/data | jq -r '.message // .error'; \
	done

health: ## Check health of all instances
	@echo "Checking health..."
	@echo -n "App 1: "; curl -s http://localhost:8081/health | jq -r '.status'
	@echo -n "App 2: "; curl -s http://localhost:8082/health | jq -r '.status'
	@echo -n "App 3: "; curl -s http://localhost:8083/health | jq -r '.status'

redis-cli: ## Connect to Redis CLI
	docker-compose exec redis redis-cli

redis-keys: ## Show rate limit keys in Redis
	@echo "Rate limit keys in Redis:"
	docker-compose exec redis redis-cli --scan --pattern "example:ratelimit:*"

redis-monitor: ## Monitor Redis commands in real-time
	docker-compose exec redis redis-cli MONITOR

clean: stop ## Stop services and clean up volumes
	@echo "Cleaning up..."
	docker-compose down -v
	@echo "Cleanup complete"

build: ## Build Docker images
	docker-compose build

run-local: ## Run locally (requires Redis on localhost:6379)
	@echo "Starting server locally..."
	@echo "Make sure Redis is running on localhost:6379"
	go run main.go
