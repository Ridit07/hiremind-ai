.PHONY: help build build-prod run stop clean logs compose-up compose-down compose-logs compose-rebuild docker-push version

# Variables
SERVICE_NAME=gateway-service
DOCKER_REGISTRY=docker.io
DOCKER_IMAGE=$(DOCKER_REGISTRY)/hiremind/$(SERVICE_NAME)
VERSION?=latest
BUILD_TIME=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
GIT_COMMIT=$(shell git rev-parse --short HEAD)
PORT=8080

# SSH key for pulling the private hiremind-proto-contracts dependency at build time.
# Override with: make build SSH_KEY=~/.ssh/some_other_key
SSH_KEY?=$(HOME)/.ssh/id_rsa
SSH_PRIVATE_KEY=$(shell cat $(SSH_KEY))

help: ## Display this help message
	@echo "HireMind Gateway Service - Docker & Development Commands"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Docker Commands
build: ## Build Docker image
	docker build \
		-t $(DOCKER_IMAGE):$(VERSION) \
		-t $(DOCKER_IMAGE):latest \
		--build-arg SSH_PRIVATE_KEY="$(SSH_PRIVATE_KEY)" \
		-f Dockerfile .

build-prod: ## Build production Docker image
	docker build \
		-t $(DOCKER_IMAGE):$(VERSION)-prod \
		--build-arg VERSION=$(VERSION) \
		--build-arg BUILD_TIME=$(BUILD_TIME) \
		--build-arg GIT_COMMIT=$(GIT_COMMIT) \
		--build-arg SSH_PRIVATE_KEY="$(SSH_PRIVATE_KEY)" \
		-f Dockerfile.prod .

run: ## Run Docker container
	docker run -it --rm \
		--env-file .env \
		-p $(PORT):$(PORT) \
		$(DOCKER_IMAGE):$(VERSION)

stop: ## Stop the running container
	docker stop hiremind-gateway-service || true

docker-push: ## Push Docker image to registry
	docker push $(DOCKER_IMAGE):$(VERSION)
	docker push $(DOCKER_IMAGE):latest

# Docker Compose Commands
compose-up: ## Start services with docker-compose
	SSH_PRIVATE_KEY="$(SSH_PRIVATE_KEY)" docker-compose up -d

compose-down: ## Stop services with docker-compose
	docker-compose down

compose-logs: ## View docker-compose logs
	docker-compose logs -f

compose-logs-service: ## View gateway-service logs only
	docker-compose logs -f gateway-service

compose-ps: ## Show running containers
	docker-compose ps

compose-rebuild: ## Rebuild container and start
	SSH_PRIVATE_KEY="$(SSH_PRIVATE_KEY)" docker-compose up -d --build

logs: ## Follow gateway-service container logs
	docker logs -f hiremind-gateway-service

# Development Commands
dev: ## Run service in development mode
	go run main.go

test: ## Run tests
	go test -v ./...

test-coverage: ## Run tests with coverage
	go test -cover ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

lint: ## Run linter
	golangci-lint run

fmt: ## Format code
	go fmt ./...

clean: ## Clean build artifacts and dangling images
	rm -f $(SERVICE_NAME)
	docker image prune -f

# Utility Commands
version: ## Display build information
	@echo "Service: $(SERVICE_NAME)"
	@echo "Version: $(VERSION)"
	@echo "Docker Image: $(DOCKER_IMAGE)"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Git Commit: $(GIT_COMMIT)"

health-check: ## Check service health
	curl -f http://localhost:$(PORT)/health || echo "Service not healthy"

image-size: ## Show Docker image size
	docker images $(DOCKER_IMAGE)

deps: ## Download Go dependencies
	go mod download
	go mod tidy

.DEFAULT_GOAL := help
