# backend/Dockerfile
# Build stage
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /coffee ./cmd/server

# Runtime stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /coffee-core /app/coffee
COPY --from=builder /app/migrations /app/migrations
EXPOSE 8080
CMD ["/app/coffee"]