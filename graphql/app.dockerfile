# ------------------------------------------------------------
# Build stage
# ------------------------------------------------------------
FROM golang:1.13-alpine3.11 AS build

# Install build deps (kept minimal)
RUN apk add --no-cache ca-certificates git

# Enable modules & vendor usage; disable CGO for a static binary
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOFLAGS="-mod=vendor"

# Work inside module path
WORKDIR /go/src/github.com/meharifitih/go-grpc-graphql-microservice

# Copy only files that affect dependency resolution first (build cache friendly)
COPY go.mod go.sum ./
COPY vendor/ vendor/

# Then copy the actual code
COPY account/ account/
COPY catalog/ catalog/
COPY order/ order/
COPY graphql/ graphql/

# Build the service (strip symbol table & DWARF, name the binary 'app')
RUN go build -ldflags="-s -w" -o /go/bin/app ./graphql

# ------------------------------------------------------------
# Runtime stage (distroless-like Alpine, non-root)
# ------------------------------------------------------------
FROM alpine:3.11

# Add non-root user/group
RUN addgroup -S app && adduser -S -G app app

# Minimal runtime deps
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy statically linked binary
COPY --from=build /go/bin/app /app/app

# Drop privileges
USER app:app

# App port
EXPOSE 8080

# (Optional) simple TCP healthcheck; adjust path if you have one
# HEALTHCHECK --interval=30s --timeout=3s --start-period=10s \
#   CMD wget -qO- http://127.0.0.1:8080/health || exit 1

CMD ["/app/app"]
