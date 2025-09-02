# Docker Setup for Go gRPC GraphQL Microservice

## Prerequisites

- Docker Engine 20.10+
- Docker Compose 2.0+
- At least 4GB RAM available for containers

## Quick Start

```bash
# Build all images
make build

# Start all services
make up

# Check status
make status

# View logs
make logs
```

## Services

### Core Services
- **account**: User account management (gRPC + PostgreSQL)
- **catalog**: Product catalog (gRPC + Elasticsearch)
- **order**: Order management (gRPC + PostgreSQL)
- **graphql**: GraphQL gateway (exposed on port 8000)

### Databases
- **account_db**: PostgreSQL 16 for account data
- **catalog_db**: Elasticsearch 8.11.3 for catalog data
- **order_db**: PostgreSQL 16 for order data

## Ports

- GraphQL Gateway: `8000:8080`
- All other services: Internal only (8080)

## Environment Variables

Key environment variables are configured in `docker-compose.yaml`:
- Database connection strings
- Service URLs for inter-service communication
- Database credentials (default: mehari/123456)

## Health Checks

All services include health checks:
- gRPC services use `grpc_health_probe`
- PostgreSQL uses `pg_isready`
- Elasticsearch uses cluster health endpoint

## Volumes

Data persistence is configured for:
- `account_data`: Account PostgreSQL data
- `order_data`: Order PostgreSQL data
- `elasticsearch_data`: Elasticsearch data

## Networks

All services run on a custom bridge network `microservice-network` with subnet `172.20.0.0/16`.

## Common Commands

```bash
# Rebuild everything
make rebuild

# Stop all services
make down

# Clean up everything
make clean

# Start individual services
make account-up
make catalog-up
make order-up
make graphql-up
```

## Troubleshooting

### Service won't start
1. Check logs: `make logs`
2. Verify database health: `docker-compose ps`
3. Check resource usage: `docker stats`

### Database connection issues
1. Ensure databases are healthy: `docker-compose ps`
2. Check database logs: `docker-compose logs account_db`
3. Verify network connectivity: `docker network inspect go-grpc-graphql-microservice_microservice-network`

### Build issues
1. Clean and rebuild: `make rebuild`
2. Check Docker daemon logs
3. Verify sufficient disk space

## Development

For development, you can:
1. Use `make build` to rebuild after code changes
2. Use `make logs` to monitor service logs
3. Use individual service commands for targeted testing
