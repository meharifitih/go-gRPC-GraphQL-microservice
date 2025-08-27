# Go gRPC GraphQL Microservice

This project is a microservice architecture using Go, gRPC, GraphQL, and PostgreSQL. It demonstrates how to build scalable services with gRPC APIs and expose them via GraphQL.

## Project Structure

- `account/` - Account service (gRPC server, proto, DB, etc.)
- `catalog/` - Catalog service (structure similar to account)
- `graphql/` - GraphQL gateway (resolvers, schema, etc.)
- `order/` - Order service (structure similar to account)
- `docker-compose.yaml` - Multi-service orchestration

## Prerequisites

- Go 1.18+
- Docker & Docker Compose
- [protoc](https://grpc.io/docs/protoc-installation/) (Protocol Buffers compiler)
- [protoc-gen-go](https://github.com/protocolbuffers/protobuf-go)
- [protoc-gen-go-grpc](https://github.com/grpc/grpc-go)
- [gqlgen](https://gqlgen.com/)

## Setup

### 1. Clone the repository

```bash
git clone https://github.com/meharifitih/go-grpc-graphql-microservice.git
cd go-grpc-graphql-microservice
```

### 2. Install Go dependencies

```bash
go mod tidy
```

### 3. Generate gRPC code from proto files

For the account service:

```bash
protoc --go_out=account/pb --go-grpc_out=account/pb account/account.proto
```

Repeat for other services as needed (e.g., `catalog.proto`, `order.proto`).

### 4. Generate GraphQL code

From the `graphql/` directory:

```bash
cd graphql
gqlgen generate
cd ..
```

### 5. Run with Docker Compose

```bash
docker-compose up --build
```

This will start all services and databases.

### 6. Run services locally (optional)

You can run each service locally for development. For example, to run the account service:

```bash
go run account/cmd/account/main.go
```

## Useful Commands

- Generate proto (account):
  ```bash
  protoc --go_out=account/pb --go-grpc_out=account/pb account/account.proto
  ```
- Generate GraphQL code:
  ```bash
  cd graphql && gqlgen generate && cd ..
  ```
- Start all services:
  ```bash
  docker-compose up --build
  ```

## Troubleshooting

- **Postgres SSL error:** If you see `pq: SSL is not enabled on the server`, ensure your DB connection string includes `sslmode=disable`.
- **Proto generation errors:** Make sure `protoc`, `protoc-gen-go`, and `protoc-gen-go-grpc` are installed and in your PATH.

## License

MIT
