# Go gRPC GraphQL Microservice

A modern, scalable microservice architecture built with Go, featuring gRPC communication between services, GraphQL API gateway, and robust data persistence. This project demonstrates enterprise-grade microservice patterns with proper separation of concerns, service discovery, and API aggregation.

## 🏗️ Architecture Overview

This project implements a **microservice architecture** with the following components:

- **Account Service**: User account management with PostgreSQL
- **Catalog Service**: Product catalog with Elasticsearch
- **Order Service**: Order processing and management with PostgreSQL
- **GraphQL Gateway**: Unified API layer aggregating all services

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Account       │    │   Catalog       │    │   Order         │
│   Service       │    │   Service       │    │   Service       │
│   (gRPC)       │    │   (gRPC)        │    │   (gRPC)        │
│   Port: 8080   │    │   Port: 8080    │    │   Port: 8080    │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         │                       │                       │
         └───────────────────────┼───────────────────────┘
                                 │
                    ┌─────────────────┐
                    │   GraphQL       │
                    │   Gateway       │
                    │   Port: 8000    │
                    └─────────────────┘
```

## 🚀 Tech Stack

### **Backend Services**
- **Go 1.24** - High-performance, compiled language
- **gRPC** - High-performance RPC framework using Protocol Buffers
- **GraphQL** - Query language and runtime for APIs
- **gqlgen** - Go GraphQL code generation

### **Databases**
- **PostgreSQL 16** - Relational database for account and order data
- **Elasticsearch 8.11.3** - Search and analytics engine for catalog data

### **Infrastructure**
- **Docker & Docker Compose** - Containerization and orchestration
- **Alpine Linux** - Lightweight container base images
- **Custom networking** - Isolated service communication

### **Development Tools**
- **Protocol Buffers** - Data serialization format
- **Make** - Build automation and common commands
- **Health checks** - Service monitoring and dependency management

## 📁 Project Structure

```
go-grpc-graphql-microservice/
├── account/                 # Account microservice
│   ├── cmd/account/        # Main application entry point
│   ├── pb/                 # Generated Protocol Buffer code
│   ├── account.proto       # Service definition
│   ├── server.go           # gRPC server implementation
│   ├── service.go          # Business logic
│   ├── repository.go       # Data access layer
│   ├── client.go           # gRPC client for inter-service calls
│   ├── app.dockerfile      # Service container definition
│   ├── db.dockerfile       # Database container definition
│   └── up.sql              # Database initialization script
├── catalog/                 # Catalog microservice
│   ├── cmd/catalog/        # Main application entry point
│   ├── pb/                 # Generated Protocol Buffer code
│   ├── catalog.proto       # Service definition
│   ├── server.go           # gRPC server implementation
│   ├── service.go          # Business logic
│   ├── repository.go       # Data access layer (Elasticsearch)
│   ├── client.go           # gRPC client for inter-service calls
│   └── app.dockerfile      # Service container definition
├── order/                   # Order microservice
│   ├── cmd/order/          # Main application entry point
│   ├── pb/                 # Generated Protocol Buffer code
│   ├── order.proto         # Service definition
│   ├── server.go           # gRPC server implementation
│   ├── service.go          # Business logic
│   ├── repository.go       # Data access layer
│   ├── client.go           # gRPC client for inter-service calls
│   ├── app.dockerfile      # Service container definition
│   ├── db.dockerfile       # Database container definition
│   └── up.sql              # Database initialization script
├── graphql/                 # GraphQL API Gateway
│   ├── cmd/graphql/        # Main application entry point
│   ├── schema.graphql      # GraphQL schema definition
│   ├── gqlgen.yml          # Code generation configuration
│   ├── generated.go         # Auto-generated GraphQL code
│   ├── resolver.go          # GraphQL resolver implementation
│   ├── app.dockerfile      # Gateway container definition
│   └── main.go             # Gateway entry point
├── docker-compose.yaml      # Multi-service orchestration
├── Makefile                 # Build and deployment commands
├── .dockerignore            # Docker build optimization
├── go.mod                   # Go module dependencies
└── README.md                # This file
```

## 🛠️ Prerequisites

### **System Requirements**
- **Docker Engine** 20.10+ with Docker Compose 2.0+
- **Go 1.24+** (for local development)
- **4GB+ RAM** available for containers
- **Git** for version control

### **Development Tools** (Optional)
- **protoc** - Protocol Buffers compiler
- **protoc-gen-go** - Go Protocol Buffer plugin
- **protoc-gen-go-grpc** - Go gRPC plugin
- **gqlgen** - GraphQL code generation

## 🚀 Quick Start

### **1. Clone the Repository**
```bash
git clone https://github.com/meharifitih/go-grpc-graphql-microservice.git
cd go-grpc-graphql-microservice
```

### **2. Start All Services**
```bash
# Build and start all services
make build
make up

# Or use Docker Compose directly
docker-compose up -d
```

### **3. Verify Services**
```bash
# Check service status
make status

# View logs
make logs

# Check individual service logs
docker-compose logs account
docker-compose logs catalog
docker-compose logs order
docker-compose logs graphql
```

### **4. Access the Application**
- **GraphQL Playground**: http://localhost:8000/playground
- **GraphQL Endpoint**: http://localhost:8000/graphql
- **Service Health**: Check with `docker-compose ps`

## 📚 Usage Guide

### **GraphQL API**

The GraphQL gateway provides a unified interface to all microservices:

```graphql
# Query accounts
query {
  accounts(skip: 0, take: 10) {
    id
    name
  }
}

# Query products
query {
  products(skip: 0, take: 10) {
    id
    name
    description
    price
  }
}

# Query orders for an account
query {
  ordersForAccount(accountId: "account_id_here") {
    id
    accountId
    totalPrice
    products {
      id
      name
      quantity
      price
    }
  }
}

# Create an account
mutation {
  postAccount(name: "John Doe") {
    id
    name
  }
}

# Create an order
mutation {
  postOrder(
    accountId: "account_id"
    products: [
      { productId: "product_id", quantity: 2 }
    ]
  ) {
    id
    totalPrice
  }
}
```

### **Service Communication**

Services communicate via gRPC:
- **Account Service**: Manages user accounts and authentication
- **Catalog Service**: Handles product catalog and search
- **Order Service**: Processes orders and manages order lifecycle
- **GraphQL Gateway**: Aggregates all services into a single API

## 🛠️ Development

### **Local Development Setup**

```bash
# Install Go dependencies
go mod tidy

# Generate Protocol Buffer code
protoc --go_out=account/pb --go-grpc_out=account/pb account/account.proto
protoc --go_out=catalog/pb --go-grpc_out=catalog/pb catalog/catalog.proto
protoc --go_out=order/pb --go-grpc_out=order/pb order/order.proto

# Generate GraphQL code
cd graphql && gqlgen generate && cd ..

# Run services locally
go run account/cmd/account/main.go
go run catalog/cmd/catalog/main.go
go run order/cmd/order/main.go
go run graphql/cmd/graphql/main.go
```

### **Docker Development**

```bash
# Rebuild specific service
docker-compose build account

# Restart specific service
docker-compose restart account

# View service logs
docker-compose logs -f account

# Access service shell
docker-compose exec account sh
```

## 📊 Make Commands

The project includes a comprehensive Makefile for common operations:

```bash
# Build all images
make build

# Start all services
make up

# Stop all services
make down

# View service status
make status

# View logs
make logs

# Clean up everything
make clean

# Rebuild and restart
make rebuild

# Individual service management
make account-up
make catalog-up
make order-up
make graphql-up
```

## 🔧 Configuration

### **Environment Variables**

Key configuration is handled via environment variables:

```yaml
# Account Service
DATABASE_URL: postgres://mehari:123456@account_db:5432/mehari?sslmode=disable

# Catalog Service
DATABASE_URL: http://catalog_db:9200

# Order Service
DATABASE_URL: postgres://mehari:123456@order_db:5432/mehari?sslmode=disable
ACCOUNT_SERVICE_URL: account:8080
CATALOG_SERVICE_URL: catalog:8080

# GraphQL Gateway
ACCOUNT_SERVICE_URL: account:8080
CATALOG_SERVICE_URL: catalog:8080
ORDER_SERVICE_URL: order:8080
```

### **Database Configuration**

- **PostgreSQL**: Default credentials (mehari/123456)
- **Elasticsearch**: Single-node setup with security disabled for development
- **Data Persistence**: Named volumes for all databases

## 🚨 Troubleshooting

### **Common Issues**

1. **Service Won't Start**
   ```bash
   # Check logs
   make logs
   
   # Verify database health
   docker-compose ps
   ```

2. **Database Connection Issues**
   ```bash
   # Check database logs
   docker-compose logs account_db
   docker-compose logs catalog_db
   docker-compose logs order_db
   ```

3. **Port Conflicts**
   ```bash
   # Check what's using port 8000
   lsof -i :8000
   
   # Stop conflicting services
   docker-compose down
   ```

4. **Memory Issues**
   ```bash
   # Check container resource usage
   docker stats
   
   # Increase Docker memory limit
   ```

### **Reset Everything**

```bash
# Complete cleanup
make clean

# Rebuild from scratch
make rebuild
```

## 🧪 Testing

### **Service Testing**

```bash
# Test individual services
curl -X POST http://localhost:8000/graphql \
  -H "Content-Type: application/json" \
  -d '{"query":"{ accounts(skip: 0, take: 5) { id name } }"}'
```

### **Integration Testing**

The services are designed to work together. Test the complete flow:
1. Create an account
2. Add products to catalog
3. Create an order
4. Query order details

## 📈 Monitoring & Observability

- **Service Health**: Built-in health checks and status monitoring
- **Logs**: Structured logging for all services
- **Metrics**: Service performance and resource usage
- **Network**: Isolated service communication with custom subnet

## 🔒 Security Considerations

- **Development Mode**: Security features disabled for local development
- **Production Ready**: Includes security configurations for production deployment
- **Network Isolation**: Services communicate only through defined interfaces
- **Data Encryption**: Support for encrypted database connections

## 🚀 Deployment

### **Production Considerations**

1. **Enable Security**: Enable Elasticsearch security features
2. **SSL/TLS**: Configure proper SSL certificates
3. **Authentication**: Implement proper service authentication
4. **Monitoring**: Add production monitoring and alerting
5. **Backup**: Implement database backup strategies
6. **Scaling**: Configure horizontal scaling for services

### **Kubernetes Deployment**

The services can be easily deployed to Kubernetes:
- Use the Docker images as base
- Configure service discovery
- Implement proper health checks
- Add resource limits and requests

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- **gRPC**: High-performance RPC framework
- **GraphQL**: Query language for APIs
- **gqlgen**: Go GraphQL code generation
- **Docker**: Containerization platform
- **Go Community**: Excellent Go ecosystem and tools

---

**Happy Coding! 🚀**

For questions or issues, please check the troubleshooting section or open an issue on GitHub.
