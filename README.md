# Go Microservices

A modern microservices architecture built with Go, featuring gRPC communication, REST API gateway, and comprehensive observability.

## 🏗️ Architecture

This project implements a microservices architecture with the following components:

- **API Gateway**: REST API gateway that routes requests to appropriate microservices
- **Auth Service**: Handles user authentication and authorization
- **User Service**: Manages user profiles and information
- **Post Service**: Manages blog posts and content
- **Shared Infrastructure**: Database, caching, monitoring, and observability

## 📁 Project Structure

```
go-microservices/
├── api-gateway/              # REST API Gateway
│   ├── cmd/main.go          # Entry point
│   ├── internal/
│   │   ├── handlers/        # HTTP handlers
│   │   ├── clients/         # gRPC clients
│   │   ├── middlewares/     # HTTP middlewares
│   │   └── routes/          # Route definitions
│   ├── config/              # Configuration
│   └── go.mod
│
├── services/
│   ├── auth-service/        # Authentication service
│   ├── user-service/        # User management service
│   └── post-service/        # Post management service
│
├── proto/                   # Protocol buffer definitions
│   ├── auth.proto
│   ├── user.proto
│   ├── post.proto
│   └── common/
│       └── types.proto
│
├── docker-compose.yml       # Development environment
├── Makefile                # Build and deployment commands
└── README.md
```

## 🚀 Quick Start

### Prerequisites

- Go 1.21+
- Docker and Docker Compose
- Protocol Buffers compiler (`protoc`)
- Make

### Development Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd go-microservices
   ```

2. **Setup development environment**
   ```bash
   make dev-setup
   ```

3. **Start the infrastructure (databases, monitoring)**
   ```bash
   make dev-start
   ```

4. **Generate protobuf files**
   ```bash
   make proto
   ```

5. **Run services locally** (in separate terminals)
   ```bash
   make run-auth     # Terminal 1
   make run-user     # Terminal 2
   make run-post     # Terminal 3
   make run-gateway  # Terminal 4
   ```

### Docker Development

For a complete Docker-based development environment:

```bash
# Build and start all services
make docker-build
make docker-up

# View logs
make docker-logs

# Stop services
make docker-down
```

## 🔧 Available Commands

### Build Commands
- `make build` - Build all services
- `make build-auth` - Build auth service
- `make build-user` - Build user service
- `make build-post` - Build post service
- `make build-gateway` - Build API gateway

### Development Commands
- `make dev-setup` - Setup development environment
- `make dev-start` - Start development infrastructure
- `make dev-stop` - Stop development environment
- `make proto` - Generate protobuf files
- `make deps` - Install dependencies

### Docker Commands
- `make docker-build` - Build Docker images
- `make docker-up` - Start all services
- `make docker-down` - Stop all services
- `make docker-logs` - View service logs
- `make docker-clean` - Clean up Docker resources

### Testing Commands
- `make test` - Run all tests
- `make test-coverage` - Run tests with coverage
- `make lint` - Run linter
- `make format` - Format code

### Database Commands
- `make db-up` - Start database services
- `make db-down` - Stop database services
- `make db-migrate` - Run database migrations

### Monitoring Commands
- `make monitoring-up` - Start monitoring stack
- `make monitoring-down` - Stop monitoring stack

## 🌐 Service Endpoints

### API Gateway (Port 8080)
- `GET /health` - Health check
- `POST /api/auth/login` - User login
- `POST /api/auth/register` - User registration
- `GET /api/users/:id` - Get user profile
- `PUT /api/users/:id` - Update user profile
- `GET /api/posts` - List posts
- `POST /api/posts` - Create post
- `GET /api/posts/:id` - Get post
- `PUT /api/posts/:id` - Update post
- `DELETE /api/posts/:id` - Delete post

### gRPC Services
- **Auth Service**: `localhost:50051`
- **User Service**: `localhost:50052`
- **Post Service**: `localhost:50053`

## 🏗️ Technology Stack

### Backend
- **Go 1.21** - Primary programming language
- **gRPC** - Inter-service communication
- **Protocol Buffers** - Message serialization
- **Fiber/Echo** - HTTP framework for API Gateway
- **GORM** - ORM for database operations
- **JWT** - Authentication tokens

### Infrastructure
- **PostgreSQL** - Primary database
- **Redis** - Caching and session storage
- **Docker** - Containerization
- **Docker Compose** - Local development orchestration

### Observability
- **Jaeger** - Distributed tracing
- **Prometheus** - Metrics collection
- **Grafana** - Metrics visualization

## 🔒 Security Features

- JWT-based authentication
- Password hashing with bcrypt
- Request rate limiting
- CORS configuration
- Input validation and sanitization

## 📊 Monitoring and Observability

### Metrics (Prometheus)
- HTTP request metrics
- gRPC call metrics
- Database connection pool metrics
- Custom business metrics

### Tracing (Jaeger)
- Distributed request tracing
- Service dependency mapping
- Performance bottleneck identification

### Dashboards (Grafana)
- Service health dashboards
- Performance metrics visualization
- Real-time monitoring alerts

Access monitoring tools:
- **Grafana**: http://localhost:3000 (admin/admin)
- **Prometheus**: http://localhost:9090
- **Jaeger**: http://localhost:16686

## 🧪 Testing

### Unit Tests
```bash
make test
```

### Integration Tests
```bash
# Start test environment
make docker-up

# Run integration tests
go test ./tests/integration/... -v
```

### Load Testing
```bash
# Example with wrk
wrk -t12 -c400 -d30s http://localhost:8080/api/posts
```

## 🚀 Deployment

### Production Deployment
```bash
make deploy
```

### Environment Variables

#### Auth Service
- `DB_HOST` - Database host
- `DB_PORT` - Database port
- `DB_USER` - Database username
- `DB_PASSWORD` - Database password
- `DB_NAME` - Database name
- `JWT_SECRET` - JWT signing secret
- `REDIS_HOST` - Redis host
- `REDIS_PORT` - Redis port

#### API Gateway
- `AUTH_SERVICE_HOST` - Auth service host
- `USER_SERVICE_HOST` - User service host
- `POST_SERVICE_HOST` - Post service host
- `JWT_SECRET` - JWT verification secret

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go best practices and idioms
- Write comprehensive tests
- Update documentation for new features
- Use conventional commit messages
- Ensure all linting passes

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Troubleshooting

### Common Issues

1. **Port already in use**
   ```bash
   make docker-down
   docker system prune -f
   ```

2. **Protocol buffer compilation errors**
   ```bash
   # Install protoc and Go plugins
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

3. **Database connection issues**
   ```bash
   # Restart database services
   make db-down
   make db-up
   ```

### Logs and Debugging

```bash
# View service logs
make docker-logs

# View specific service logs
docker-compose logs -f auth-service

# Debug mode (if implemented)
DEBUG=true make run-auth
```

## 📞 Support

For questions and support:
- Create an issue in the repository
- Check existing documentation
- Review the troubleshooting section

---

**Happy coding! 🚀**
