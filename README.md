# Go Microservices

A distributed microservices architecture built with Go, featuring authentication, logging, messaging, and email services with Docker containerization and RabbitMQ event-driven communication.

## Architecture Overview

This project demonstrates a complete microservices ecosystem with the following services:

- **Broker Service** - API Gateway and service orchestrator
- **Authentication Service** - User authentication and authorization
- **Logger Service** - Centralized logging with MongoDB
- **Mail Service** - Email notifications with templates
- **Listener Service** - Event processing and message handling
- **Front-end Service** - Web interface

## Services

### 🔗 Broker Service
- **Port**: 4001
- **Purpose**: Acts as an API gateway, routing requests to appropriate services
- **Features**: CORS handling, request routing, service communication
- **Tech**: Go, Chi router, gRPC, RabbitMQ

### 🔐 Authentication Service
- **Port**: 4002
- **Purpose**: Handles user authentication and authorization
- **Database**: PostgreSQL
- **Features**: User management, JWT tokens, secure authentication

### 📝 Logger Service
- **Purpose**: Centralized logging system
- **Database**: MongoDB
- **Features**: Log aggregation, gRPC communication, structured logging
- **Protocol**: gRPC with Protocol Buffers

### 📧 Mail Service
- **Purpose**: Email notifications and messaging
- **Features**: HTML/Plain text templates, SMTP integration
- **Templates**: Customizable email templates
- **Integration**: MailHog for development testing

### 👂 Listener Service
- **Purpose**: Event-driven message processing
- **Features**: RabbitMQ consumer, event handling, asynchronous processing

### 🌐 Front-end Service
- **Purpose**: Web interface for the microservices
- **Tech**: Go templates, HTTP server

## Prerequisites

- Go 1.23+
- Docker & Docker Compose
- Make (for build automation)

## Quick Start

### 1. Clone the Repository
```bash
git clone <repository-url>
cd go-microservices
```

### 2. Start All Services
```bash
cd project
make up_build
```

This will:
- Build all service binaries
- Start Docker containers
- Initialize databases
- Set up message queues

### 3. Access Services
- **Broker Service**: http://localhost:4001
- **Authentication Service**: http://localhost:4002
- **MailHog UI**: http://localhost:8025
- **PostgreSQL**: localhost:5432
- **MongoDB**: localhost:27017
- **RabbitMQ**: localhost:5672

## Development

### Build Individual Services
```bash
# Build all services
make up_build

# Build specific services
make build_broker
make build_auth
make build_logger
make build_mail
make build_listener
```

### Start Front-end Only
```bash
make build_front
make start
```

### Stop Services
```bash
make down
make stop
```

## Infrastructure

### Databases
- **PostgreSQL**: User data for authentication service
- **MongoDB**: Log storage for logger service

### Message Queue
- **RabbitMQ**: Event-driven communication between services

### Reverse Proxy
- **Caddy**: Load balancing and reverse proxy (configured via Caddyfile)

### Development Tools
- **MailHog**: Email testing and debugging

## Environment Variables

### Mail Service
```env
MAIL_DOMAIN=localhost
MAIL_HOST=mailhog
MAIL_PORT=1025
MAIL_ENCRYPTION=none
FROM_NAME="John Smith"
FROM_ADDRESS=john.smith@example.com
```

### Authentication Service
```env
DSN="host=postgres port=5432 user=postgres password=password dbname=users sslmode=disable timezone=UTC connect_timeout=5"
```

## Project Structure

```
├── authentication-service/    # User authentication & authorization
├── broker-service/           # API gateway & service orchestration
├── front-end/               # Web interface
├── listener-service/        # Event processing & message handling
├── logger-service/          # Centralized logging system
├── mail-service/           # Email notifications
└── project/               # Docker configuration & deployment
    ├── docker-compose.yml # Service orchestration
    ├── Makefile          # Build automation
    └── Caddyfile         # Reverse proxy config
```

## Communication Patterns

- **HTTP/REST**: Client-to-service and service-to-service communication
- **gRPC**: High-performance communication (Logger service)
- **RabbitMQ**: Asynchronous event-driven messaging
- **Protocol Buffers**: Efficient serialization for gRPC

## Development Features

- **Hot Reload**: Automatic rebuilding during development
- **Containerization**: Full Docker support for all services
- **Service Discovery**: Automatic service registration and discovery
- **Health Checks**: Built-in health monitoring
- **Logging**: Centralized structured logging
- **Email Testing**: MailHog integration for email debugging

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

This project is for educational purposes demonstrating microservices architecture patterns in Go.
