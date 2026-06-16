#!/bin/bash
# Docker build script for Consumer Service

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Configuration
SERVICE_NAME="consumer-service"
DOCKER_REGISTRY="${DOCKER_REGISTRY:-docker.io}"
DOCKER_IMAGE="${DOCKER_REGISTRY}/hiremind/${SERVICE_NAME}"
VERSION="${VERSION:-latest}"
BUILD_TIME=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
GIT_COMMIT=$(git rev-parse --short HEAD)
DOCKERFILE="${DOCKERFILE:-Dockerfile}"
PUSH="${PUSH:-false}"
NO_CACHE="${NO_CACHE:-false}"

# Help function
show_help() {
    cat << EOF
Usage: $0 [OPTIONS]

Docker build script for $SERVICE_NAME

OPTIONS:
    -v, --version VERSION       Docker image version (default: latest)
    -r, --registry REGISTRY     Docker registry (default: docker.io)
    -f, --file DOCKERFILE       Dockerfile to use (default: Dockerfile)
    -p, --push                  Push image to registry after building
    --no-cache                  Build without using cache
    -h, --help                  Show this help message

EXAMPLES:
    $0                          # Build with default settings
    $0 -v 1.0.0                 # Build version 1.0.0
    $0 -v 1.0.0 --push          # Build and push version 1.0.0
    $0 -f Dockerfile.prod -p    # Build production image and push

EOF
}

# Parse arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        -v|--version)
            VERSION="$2"
            shift 2
            ;;
        -r|--registry)
            DOCKER_REGISTRY="$2"
            DOCKER_IMAGE="${DOCKER_REGISTRY}/hiremind/${SERVICE_NAME}"
            shift 2
            ;;
        -f|--file)
            DOCKERFILE="$2"
            shift 2
            ;;
        -p|--push)
            PUSH="true"
            shift
            ;;
        --no-cache)
            NO_CACHE="true"
            shift
            ;;
        -h|--help)
            show_help
            exit 0
            ;;
        *)
            echo -e "${RED}Unknown option: $1${NC}"
            show_help
            exit 1
            ;;
    esac
done

# Validate Dockerfile exists
if [ ! -f "$DOCKERFILE" ]; then
    echo -e "${RED}Error: Dockerfile not found: $DOCKERFILE${NC}"
    exit 1
fi

# Build image
echo -e "${YELLOW}Building Docker image...${NC}"
echo "Image: ${DOCKER_IMAGE}:${VERSION}"
echo "Dockerfile: $DOCKERFILE"
echo "Build Time: $BUILD_TIME"
echo "Git Commit: $GIT_COMMIT"
echo ""

BUILD_ARGS="--build-arg VERSION=${VERSION} --build-arg BUILD_TIME=${BUILD_TIME} --build-arg GIT_COMMIT=${GIT_COMMIT}"

if [ "$NO_CACHE" = "true" ]; then
    BUILD_ARGS="$BUILD_ARGS --no-cache"
fi

docker build \
    $BUILD_ARGS \
    -t "${DOCKER_IMAGE}:${VERSION}" \
    -t "${DOCKER_IMAGE}:latest" \
    -f "$DOCKERFILE" \
    .

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Build successful!${NC}"
    
    # Display image info
    echo ""
    echo "Image details:"
    docker images "${DOCKER_IMAGE}"
    
    # Push if requested
    if [ "$PUSH" = "true" ]; then
        echo ""
        echo -e "${YELLOW}Pushing Docker image...${NC}"
        docker push "${DOCKER_IMAGE}:${VERSION}"
        docker push "${DOCKER_IMAGE}:latest"
        
        if [ $? -eq 0 ]; then
            echo -e "${GREEN}✓ Push successful!${NC}"
        else
            echo -e "${RED}✗ Push failed!${NC}"
            exit 1
        fi
    fi
else
    echo -e "${RED}✗ Build failed!${NC}"
    exit 1
fi
