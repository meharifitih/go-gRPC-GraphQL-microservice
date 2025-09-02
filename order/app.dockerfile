FROM golang:1.24-alpine AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 go build -o /app/order ./order/cmd/order

FROM alpine:3.19
RUN apk --no-cache add ca-certificates netcat-openbsd
WORKDIR /app
COPY --from=build /app/order .
EXPOSE 8080
CMD ["./order"]