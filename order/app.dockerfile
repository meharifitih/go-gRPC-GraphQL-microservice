# ---- Build Stage ----
FROM golang:1.23-alpine AS build

# Install build dependencies
RUN apk add --no-cache build-base ca-certificates

# Set working directory
WORKDIR /app

# Copy go module files first for caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy vendor folder if you’re vendoring (optional)
COPY vendor ./vendor

# Copy your source code
COPY order ./order

# Build the binary
RUN go build -mod=vendor -o /bin/app ./order/cmd/order

# ---- Runtime Stage ----
FROM alpine:3.21

# Add CA certificates for HTTPS
RUN apk add --no-cache ca-certificates

# Set working dir
WORKDIR /usr/bin

# Copy the built binary from the build stage
COPY --from=build /bin/app .

# Expose the service port
EXPOSE 8080

# Run the app
CMD ["./app"]
