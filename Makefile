# Go Microservices Makefile

# Variables
PROTO_DIR := proto
SERVICES := auth-service user-service post-service
API_GATEWAY := api-gateway
DOCKER_COMPOSE := docker-compose.yml

# Go variables
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
GO_VERSION := 1.21

# Proto compilation
.PHONY: proto
proto:
	@echo "Generating protobuf files..."
	@for service in $(SERVICES); do \
		echo "Generating proto for $$service..."; \
		mkdir -p services/$$service/proto; \
		protoc --go_out=services/$$service/proto --go_opt=paths=source_relative \
			--go-grpc_out=services/$$service/proto --go-grpc_opt=paths=source_relative \
			--proto_path=$(PROTO_DIR) \
			$(PROTO_DIR)/*.proto $(PROTO_DIR)/common/*.proto; \
	done
	@echo "Protobuf generation completed!"

# Build individual services
.PHONY: build-auth
build-auth:
	@echo "Building auth-service..."
	cd services/auth-service && go build -o bin/auth-service ./cmd/main.go

.PHONY: build-user
build-user:
	@echo "Building user-service..."
	cd services/user-service && go build -o bin/user-service ./cmd/main.go

.PHONY: build-post
build-post:
	@echo "Building post-service..."
	cd services/post-service && go build -o bin/post-service ./cmd/main.go

.PHONY: build-gateway
build-gateway:
	@echo "Building api-gateway..."
	cd api-gateway && go build -o bin/api-gateway ./cmd/main.go

# Build all services
.PHONY: build
build: build-auth build-user build-post build-gateway
	@echo "All services built successfully!"

# Run individual services locally
.PHONY: run-auth
run-auth:
	cd services/auth-service && go run ./cmd/main.go

.PHONY: run-user
run-user:
	cd services/user-service && go run ./cmd/main.go

.PHONY: run-post
run-post:
	cd services/post-service && go run ./cmd/main.go

.PHONY: run-gateway
run-gateway:
	cd api-gateway && go run ./cmd/main.go

# Docker operations
.PHONY: docker-build
docker-build:
	@echo "Building Docker images..."
	docker-compose build

.PHONY: docker-up
docker-up:
	@echo "Starting services with Docker Compose..."
	docker-compose up -d

.PHONY: docker-down
docker-down:
	@echo "Stopping services..."
	docker-compose down

.PHONY: docker-logs
docker-logs:
	docker-compose logs -f

.PHONY: docker-clean
docker-clean:
	@echo "Cleaning up Docker resources..."
	docker-compose down -v
	docker system prune -f

# Development operations
.PHONY: deps
deps:
	@echo "Installing dependencies for all services..."
	@for service in $(SERVICES); do \
		echo "Installing deps for $$service..."; \
		cd services/$$service && go mod tidy && go mod download; \
		cd ../..; \
	done
	cd $(API_GATEWAY) && go mod tidy && go mod download

.PHONY: test
test:
	@echo "Running tests for all services..."
	@for service in $(SERVICES); do \
		echo "Testing $$service..."; \
		cd services/$$service && go test ./... -v; \
		cd ../..; \
	done
	cd $(API_GATEWAY) && go test ./... -v

.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	@for service in $(SERVICES); do \
		echo "Testing $$service with coverage..."; \
		cd services/$$service && go test ./... -coverprofile=coverage.out -v; \
		go tool cover -html=coverage.out -o coverage.html; \
		cd ../..; \
	done

.PHONY: lint
lint:
	@echo "Running linter for all services..."
	@for service in $(SERVICES); do \
		echo "Linting $$service..."; \
		cd services/$$service && golangci-lint run; \
		cd ../..; \
	done
	cd $(API_GATEWAY) && golangci-lint run

.PHONY: format
format:
	@echo "Formatting code..."
	@for service in $(SERVICES); do \
		echo "Formatting $$service..."; \
		cd services/$$service && gofmt -w .; \
		cd ../..; \
	done
	cd $(API_GATEWAY) && gofmt -w .

# Database operations
.PHONY: db-up
db-up:
	@echo "Starting database..."
	docker-compose up -d postgres redis

.PHONY: db-down
db-down:
	@echo "Stopping database..."
	docker-compose stop postgres redis

.PHONY: db-migrate
db-migrate:
	@echo "Running database migrations..."
	# Add migration commands here when implemented

# Monitoring
.PHONY: monitoring-up
monitoring-up:
	@echo "Starting monitoring stack..."
	docker-compose up -d prometheus grafana jaeger

.PHONY: monitoring-down
monitoring-down:
	@echo "Stopping monitoring stack..."
	docker-compose stop prometheus grafana jaeger

# Development workflow
.PHONY: dev-setup
dev-setup: deps proto
	@echo "Development environment setup complete!"

.PHONY: dev-start
dev-start: db-up
	@echo "Starting development environment..."
	@echo "Database is running. Start your services manually or use make run-<service>"

.PHONY: dev-stop
dev-stop: docker-down
	@echo "Development environment stopped."

# Production deployment
.PHONY: deploy
deploy: proto build docker-build docker-up
	@echo "Deployment completed!"

# Clean up
.PHONY: clean
clean:
	@echo "Cleaning up build artifacts..."
	@for service in $(SERVICES); do \
		rm -rf services/$$service/bin; \
		rm -rf services/$$service/coverage.*; \
	done
	rm -rf $(API_GATEWAY)/bin
	rm -rf $(API_GATEWAY)/coverage.*

# Help
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  proto          - Generate protobuf files"
	@echo "  build          - Build all services"
	@echo "  build-<service> - Build specific service (auth, user, post, gateway)"
	@echo "  run-<service>  - Run specific service locally"
	@echo "  docker-build   - Build Docker images"
	@echo "  docker-up      - Start all services with Docker"
	@echo "  docker-down    - Stop all services"
	@echo "  docker-logs    - View service logs"
	@echo "  docker-clean   - Clean up Docker resources"
	@echo "  deps           - Install dependencies"
	@echo "  test           - Run tests"
	@echo "  test-coverage  - Run tests with coverage"
	@echo "  lint           - Run linter"
	@echo "  format         - Format code"
	@echo "  db-up          - Start database services"
	@echo "  db-down        - Stop database services"
	@echo "  monitoring-up  - Start monitoring stack"
	@echo "  monitoring-down - Stop monitoring stack"
	@echo "  dev-setup      - Setup development environment"
	@echo "  dev-start      - Start development environment"
	@echo "  dev-stop       - Stop development environment"
	@echo "  deploy         - Deploy to production"
	@echo "  clean          - Clean up build artifacts"
	@echo "  help           - Show this help message"

# Default target
.DEFAULT_GOAL := help
