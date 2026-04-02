# Go DevSecOps REST API

A lightweight REST API application built with Go using only the standard library. Designed to be easily containerized and deployed in DevSecOps pipelines.

## Endpoints

### GET /health

Health check endpoint that returns the application status.

**Response (200 OK):**

```json
{ "status": "healthy" }
```

### GET /

Root endpoint that returns a greeting message.

**Response (200 OK):**

```json
{ "message": "Hello from DevSecOps pipeline" }
```

### GET /version

Version endpoint that returns the application version.

**Response (200 OK):**

```json
{ "version": "1.0.0" }
```

## Features

- ✅ Written in Go using only the standard library
- ✅ Configurable port via environment variable `APP_PORT` (defaults to 8080)
- ✅ Comprehensive error handling
- ✅ Multi-stage Docker build for minimal image size
- ✅ Runs as non-root user in Docker
- ✅ Includes .dockerignore for optimized builds

## Local Setup & Running

### Prerequisites

- Go 1.21 or later

### Running Locally

1. Navigate to the project directory:

```bash
cd go-devsecops-app
```

2. Run the application:

```bash
go run main.go
```

The server will start on `http://localhost:8080` by default.

3. (Optional) Specify a custom port:

```bash
APP_PORT=3000 go run main.go
```

### Testing Endpoints

```bash
# Health check
curl http://localhost:8080/health

# Root endpoint
curl http://localhost:8080/

# Version
curl http://localhost:8080/version
```

## Docker

### Building the Docker Image

```bash
docker build -t go-devsecops-app:latest .
```

### Running the Docker Container

```bash
# Using default port (8080)
docker run -p 8080:8080 go-devsecops-app:latest

# Using custom port
docker run -p 3000:8080 -e APP_PORT=8080 go-devsecops-app:latest
```

### Testing in Docker

```bash
# After running the container
curl http://localhost:8080/health
curl http://localhost:8080/
curl http://localhost:8080/version
```

## Security Features

- **Non-root user**: The Docker container runs as user ID 1000 (non-root)
- **Minimal image**: Uses multi-stage build with scratch as base, reducing attack surface
- **No external dependencies**: Reduces supply chain risk by using only Go standard library
- **CA certificates included**: Supports HTTPS operations in minimal image

## Project Structure

```
go-devsecops-app/
├── main.go          # Application source code
├── go.mod           # Go module definition
├── Dockerfile       # Multi-stage Docker build configuration
├── .dockerignore    # Docker build context optimization
└── README.md        # This file
```

## Building the Binary

```bash
# Build a statically linked binary
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app main.go

# Run the binary
./app

# Run with custom port
APP_PORT=3000 ./app
```

## Environment Variables

- `APP_PORT`: The port to listen on (default: 8080)

## License

MIT
