# Osmity-web-backend
This repository contains the backend API for **Osmity** and **Shizuku86**.  
The backend is implemented in Go and is deployed as a Docker container, with
separate configurations for development and production environments.

## Overview
- Language: Go
- Purpose:
  - Provide RESTful APIs
  - Serve Swagger (OpenAPI) documentation in development
- The backend is designed to be consumed by a Next.js frontend via `/api/*`.

## Environments
### Development (dev)
- Domains  
  - https://dev.osmity.com  
  - https://dev.shizuku86.com

- API Base: https://dev.osmity.com/api/*
- Swagger: https://dev.osmity.com/swagger/index.html
  > Swagger is **only available in the development environment**.

### Production (prod)
- Domains  
- https://osmity.com  
- https://shizuku86.com

- API Base: https://osmity.com/api/*
- Swagger: Disabled / not publicly accessible in production

## Running Locally / Development
### Using Docker
The project is expected to be structured as follows:
```
Osmity/
├── docker-compose.yml
├── osmity-web-backend/
└── osmity-web-frontend/
```
Run Docker Compose **from the project root directory** (one level above `back`):
```
cd Osmity
docker compose up -d --build
```

After startup, the services are available at:
API: http://localhost:8080/api/*
Swagger: http://localhost:8080/swagger/index.html

### Environment Variables
The backend relies on the following environment variables:
| Variable     | Description |
|-------------|-------------|
| `APP_ENV`   | Runtime environment (`dev` or `prod`). Used to control behavior such as logging level and Swagger availability. |
| `VERSION`   | Application version (e.g. semantic version or release tag). |
| `BUILD_TIME`| Build timestamp in UTC (ISO 8601 format recommended). |
| `GIT_COMMIT`| Git commit SHA used for the build (short or full). |

## Deployment
Docker images are built via GitHub Actions
Images are pushed to Docker Hub
