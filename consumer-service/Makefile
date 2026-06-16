.PHONY: help build run stop clean logs compose-up compose-down compose-logs docker-build docker-push version

# Variables
SERVICE_NAME=consumer-service
DOCKER_REGISTRY=docker.io
DOCKER_IMAGE=$(DOCKER_REGISTRY)/hiremind/$(SERVICE_NAME)
VERSION?=latest
BUILD_TIME=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
GIT_COMMIT=$(shell git rev-parse --short HEAD)

help: ## Display this help message
	@echo "HireMind Consumer Service - Docker & Development Commands"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Docker Commands
build: ## Build Docker image
	docker build \
		-t $(DOCKER_IMAGE):$(VERSION) \
		-t $(DOCKER_IMAGE):latest \
		--build-arg VERSION=$(VERSION) \
		--build-arg BUILD_TIME=$(BUILD_TIME) \
		--build-arg GIT_COMMIT=$(GIT_COMMIT) \
		-f Dockerfile .

build-prod: ## Build production Docker image
	docker build \
		-t $(DOCKER_IMAGE):$(VERSION)-prod \
		--build-arg VERSION=$(VERSION) \
		--build-arg BUILD_TIME=$(BUILD_TIME) \
		--build-arg GIT_COMMIT=$(GIT_COMMIT) \
		-f Dockerfile.prod .

run: ## Run Docker container
	docker run -it --rm \
		--env-file .env \
		-p 50051:50051 \
		$(DOCKER_IMAGE):$(VERSION)

docker-push: ## Push Docker image to registry
	docker push $(DOCKER_IMAGE):$(VERSION)
	docker push $(DOCKER_IMAGE):latest

# Docker Compose Commands
compose-up: ## Start services with docker-compose
	docker-compose up -d

compose-down: ## Stop services with docker-compose
	docker-compose down

compose-logs: ## View docker-compose logs
	docker-compose logs -f

compose-logs-service: ## View consumer-service logs only
	docker-compose logs -f consumer-service

compose-logs-db: ## View postgres logs only
	docker-compose logs -f postgres

compose-ps: ## Show running containers
	docker-compose ps

compose-rebuild: ## Rebuild container and start
	docker-compose up -d --build

# Development Commands
dev: ## Run service in development mode
	go run main.go

test: ## Run tests
	go test -v ./...

test-coverage: ## Run tests with coverage
	go test -cover ./... -out=coverage.out
	go tool cover -html=coverage.out

lint: ## Run linter
	golangci-lint run

fmt: ## Format code
	go fmt ./...

clean: ## Clean build artifacts
	rm -f $(SERVICE_NAME)
	docker image prune -f

# Database Commands
db-migrate-up: ## Run database migrations
	docker-compose exec postgres psql -U hiremind_user -d hiremind_db -c "SELECT * FROM migrations;"

db-shell: ## Open database shell
	docker-compose exec postgres psql -U hiremind_user -d hiremind_db

# Utility Commands
version: ## Display build information
	@echo "Service: $(SERVICE_NAME)"
	@echo "Version: $(VERSION)"
	@echo "Docker Image: $(DOCKER_IMAGE)"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Git Commit: $(GIT_COMMIT)"

health-check: ## Check service health
	curl -f http://localhost:50051/health || echo "Service not healthy"

logs-follow: ## Follow application logs
	tail -f logs/app.log

image-size: ## Show Docker image size
	docker images $(DOCKER_IMAGE)

deps: ## Download Go dependencies
	go mod download
	go mod tidy

.DEFAULT_GOAL := help
