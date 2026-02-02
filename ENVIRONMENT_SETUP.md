# Environment Setup Guide

This document explains how to set up and use different environments (local, production) with Docker for the CarPeek application.

## Overview

The CarPeek application supports multiple environments using environment-specific `.env` files located in the root directory. Docker Compose is configured to automatically pick up the appropriate environment file based on the `ENV` variable.

## Environment Files

The following environment files are used in the project:

### Root Directory
- `.env.local` - Used for local development
- `.env.production` - Used for production deployments

### Frontend Directory
- `frontend/.env.local` - Frontend-specific local environment variables
- `frontend/.env.production` - Frontend-specific production environment variables

Note: Backend environment variables are now consolidated in the root .env files to simplify configuration.

## Using Different Environments

### Local Environment (Default)

To run the application with local environment settings (default):

```bash
docker-compose up
```

Or explicitly specify the local environment:

```bash
ENV=local docker-compose up
```

### Production Environment

To run the application with production environment settings:

```bash
ENV=production docker-compose up
```

### Backend Only with Different Environments

For running only the backend services:

Local environment:
```bash
ENV=local docker-compose -f docker-compose-backend-only.yml up
```

Production environment:
```bash
ENV=production docker-compose -f docker-compose-backend-only.yml up
```

## Building Images for Different Environments

When building Docker images, you can specify the environment:

```bash
# Build for local environment
ENV=local docker-compose build

# Build for production environment
ENV=production docker-compose build
```

## Environment Variables Precedence

Docker Compose follows this precedence for environment variables (highest to lowest):
1. Variables set at the command line (`ENV=production docker-compose up`)
2. Variables in the environment-specific `.env` files
3. Variables set in the container
4. Default values defined in the compose files

## Adding New Environment Variables

1. Add the new variable to all relevant `.env.*` files in the root directory
2. Reference the variable in the appropriate services in the docker-compose files
3. If needed, update the application code to use the new environment variable

## Important Notes

- The frontend is a Next.js application that is exported as static files, so environment variables must be available at build time.
- The frontend Dockerfile is configured to copy the appropriate `.env` file during the build stage based on the `ENV` variable.
- Backend environment variables are now centralized in the root .env files for simplicity.
- Always ensure that sensitive information is stored securely and never committed to version control.