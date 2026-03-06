#!/bin/bash

# build-docker.sh - Build Docker images with proper environment configuration

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$SCRIPT_DIR"

usage() {
    echo "Usage: $0 [local|production]"
    echo ""
    echo "Options:"
    echo "  local      - Use .env.local for development (default)"
    echo "  production - Use .env.production for production"
    echo ""
    echo "This script:"
    echo "  1. Loads environment variables from the specified .env file"
    echo "  2. Sets up shell environment for Docker Compose"
    echo "  3. Builds Docker images"
}

# Default to local environment
ENV_TYPE="local"

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        local)
            ENV_TYPE="local"
            shift
            ;;
        production)
            ENV_TYPE="production"
            shift
            ;;
        -h|--help)
            usage
            exit 0
            ;;
        *)
            echo "Error: Unknown option '$1'"
            usage
            exit 1
            ;;
    esac
done

echo "Building with $ENV_TYPE environment..."

export ENV=$ENV_TYPE

# Determine which .env file to use
if [ "$ENV_TYPE" = "local" ]; then
    ENV_FILE="$PROJECT_ROOT/.env.local"
else
    ENV_FILE="$PROJECT_ROOT/.env.production"
fi

# Check if the .env file exists
if [ ! -f "$ENV_FILE" ]; then
    echo "Error: Environment file not found: $ENV_FILE"
    echo "Please create $ENV_FILE first."
    exit 1
fi

echo "Loading environment variables from: $ENV_FILE"

# Load environment variables from .env file into current shell
# This function safely sources the .env file, handling comments and empty lines
load_env_file() {
    local env_file="$1"
    while IFS= read -r line || [[ -n "$line" ]]; do
        # Skip empty lines and comments
        [[ -z "$line" ]] && continue
        [[ "$line" =~ ^\# ]] && continue
        
        # Export the variable
        export "$line"
    done < "$env_file"
}

# Load the environment variables
load_env_file "$ENV_FILE"

# Set ENV variable for docker-compose
export ENV="$ENV_TYPE"

# Verify required variables are set
REQUIRED_VARS=("DB_PASSWORD" "DB_DB" "DB_USER")
for var in "${REQUIRED_VARS[@]}"; do
    if [ -z "${!var:-}" ]; then
        echo "Warning: Required environment variable '$var' is not set"
        echo "Using default value for $var..."
    fi
done

# Display loaded environment variables (for debugging)
echo "Environment variables loaded:
- ENV: $ENV
- DB_PASSWORD: ${DB_PASSWORD:0:4}... (masked)
- DB_DB: $DB_DB
- DB_USER: $DB_USER
- DATABASE_URL: ${DATABASE_URL:-not set}
- NEXT_PUBLIC_API_URL: ${NEXT_PUBLIC_API_URL:-not set}"

# Clean up any existing containers (preserve volumes/data)
echo "Cleaning up existing containers..."
docker compose down 2>/dev/null || true

# Build Docker images
echo "Building Docker images..."
docker compose up -d --build
