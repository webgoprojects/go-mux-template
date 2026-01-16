# Go Mux Template

A production-ready HTTP server template using Gorilla Mux with structured logging, configuration management, middleware, and graceful shutdown.

## Features

- **HTTP Server** with Gorilla Mux router
- **Structured Logging** using Uber's Zap logger
- **Environment Configuration** management
- **Middleware** support (CORS, Request ID, HTTP Logging)
- **Health Check** endpoint for monitoring
- **Graceful Shutdown** handling
- **Template Rendering** with HTML templates
- **Static File Serving**

## Project Structure

```
go-mux-template/
├── cmd/
│   └── main.go              # Application entry point
├── pkg/
│   ├── config/
│   │   └── config.go        # Configuration management
│   ├── handlers/
│   │   ├── handlers.go      # HTTP handlers
│   │   └── health.go        # Health check handler
│   ├── logger/
│   │   └── logger.go        # Structured logging setup
│   └── middleware/
│       └── middleware.go     # HTTP middleware
├── static/                  # Static assets (CSS, JS)
├── templates/               # HTML templates
├── go.mod                   # Go module dependencies
└── README.md               # This file
```

## Getting Started

### Prerequisites

- Go 1.21 or higher

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd go-mux-template
```

2. Download dependencies:
```bash
go mod download
```

### Running

Run the server:
```bash
go run cmd/main.go
```

The server will start on `http://localhost:8080` by default.

### Building

Build the application:
```bash
go build -o bin/server cmd/main.go
```

Run the binary:
```bash
./bin/server
```

## Configuration

The application can be configured using environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Server port |
| `ENVIRONMENT` | `development` | Environment name (development/production) |
| `LOG_LEVEL` | `info` | Logging level (debug/info/warn/error) |
| `READ_TIMEOUT` | `15` | HTTP read timeout in seconds |
| `WRITE_TIMEOUT` | `15` | HTTP write timeout in seconds |

Example:
```bash
PORT=3000 ENVIRONMENT=production LOG_LEVEL=info go run cmd/main.go
```

## API Endpoints

### Web Routes

- `GET /` - Home page
- `GET /about` - About page

### API Routes

- `GET /health` - Health check endpoint

Example health check response:
```json
{
  "status": "healthy",
  "timestamp": "2024-01-01T12:00:00Z",
  "service": "go-mux-template"
}
```

## Upgrades

This template includes the following upgrades:

### Upgrade 1: Code Structure & Module Setup
- Created `go.mod` for dependency management
- Properly wired handlers package
- Fixed main.go to use handlers correctly
- Improved file path handling for templates and static files

### Upgrade 2: Structured Logging
- Integrated Uber's Zap logger for structured logging
- Added logging to all handlers
- Configurable log levels (development/production)
- Request logging with context information

### Upgrade 3: Environment Configuration
- Created config package for environment variable management
- Support for multiple configuration options
- Default values for all settings
- Easy to extend with new configuration options

### Upgrade 4: Middleware
- **Request ID Middleware**: Adds unique request ID to each request
- **CORS Middleware**: Handles Cross-Origin Resource Sharing
- **Logging Middleware**: Logs all HTTP requests with status codes and duration
- Middleware chain applied to all routes

### Upgrade 5: Health Check & Graceful Shutdown
- Health check endpoint (`/health`) for monitoring
- Graceful shutdown on SIGINT/SIGTERM signals
- Configurable server timeouts (read, write, idle)
- Proper cleanup on shutdown

## Development

### Adding New Routes

1. Create a handler in `pkg/handlers/`:
```go
func NewHandler(w http.ResponseWriter, r *http.Request) {
    // Handler logic
}
```

2. Register the route in `cmd/main.go`:
```go
r.HandleFunc("/new-route", handlers.NewHandler)
```

### Adding New Middleware

1. Create middleware function in `pkg/middleware/middleware.go`:
```go
func NewMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Middleware logic
        next.ServeHTTP(w, r)
    })
}
```

2. Register in `cmd/main.go`:
```go
r.Use(middleware.NewMiddleware)
```

## Logging

The application uses structured logging with Zap. Logs include:
- Request method and path
- Response status codes
- Request duration
- Request IDs for tracing
- Remote addresses

Example log output:
```
{"level":"info","ts":1234567890,"msg":"HTTP request","method":"GET","path":"/","status":200,"duration":0.001,"request_id":"uuid-here","remote_addr":"127.0.0.1:12345"}
```

## Graceful Shutdown

The server handles graceful shutdown on SIGINT (Ctrl+C) or SIGTERM signals:
1. Stops accepting new connections
2. Waits up to 10 seconds for existing requests to complete
3. Shuts down cleanly

## License

See LICENSE file for details.
