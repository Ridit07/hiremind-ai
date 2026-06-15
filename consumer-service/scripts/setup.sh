#!/bin/bash
# Setup script for Docker development environment

set -e

# Colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}🚀 Setting up Docker development environment...${NC}"

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo "Docker is not installed. Please install Docker Desktop."
    exit 1
fi

# Copy .env.example to .env if it doesn't exist
if [ ! -f .env ]; then
    echo -e "${BLUE}Creating .env file from .env.example...${NC}"
    cp .env.example .env
    echo -e "${GREEN}✓ .env created${NC}"
else
    echo -e "${GREEN}✓ .env already exists${NC}"
fi

# Make scripts executable
echo -e "${BLUE}Making scripts executable...${NC}"
chmod +x scripts/*.sh

# Build Docker image
echo -e "${BLUE}Building Docker image...${NC}"
make build

# Start services
echo -e "${BLUE}Starting services with docker-compose...${NC}"
make compose-up

# Wait for services to be healthy
echo -e "${BLUE}Waiting for services to be healthy...${NC}"
sleep 5

# Check service status
echo -e "${BLUE}Checking service status...${NC}"
docker-compose ps

echo -e "${GREEN}✓ Setup complete!${NC}"
echo ""
echo "Next steps:"
echo "  - View logs: make compose-logs"
echo "  - Stop services: make compose-down"
echo "  - Run tests: make test"
echo "  - View more commands: make help"
