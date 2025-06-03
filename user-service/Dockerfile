# Stage 1: Build
FROM golang:1.24.1-alpine3.20 AS builder

# Install dependencies
RUN apk update && \
    apk add --no-cache git openssh tzdata build-base python3 net-tools

WORKDIR /app

# Copy project files
COPY .env.example .env
COPY . .

# Install gin (optional CLI), tidy modules, build the app
RUN go install github.com/buu700/gin@latest
RUN go mod tidy


RUN make build

# Stage 2: Runtime
FROM alpine:latest

# Install runtime dependencies
RUN apk update && \
    apk upgrade && \
    apk add --no-cache tzdata curl && \
    mkdir -p /app

WORKDIR /app

EXPOSE 8001

# Copy built binary from builder stage
COPY --from=builder /app /app

# Start the service
ENTRYPOINT ["/app/user-service"]
