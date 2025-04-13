# Build stage
FROM golang:1.23-alpine AS builder

# Install necessary tools
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/main .
COPY .ENV .ENV

# Install PostgreSQL client and wait-for-it script
RUN apk add --no-cache postgresql-client curl bash \
    && curl -o /usr/local/bin/wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh \
    && chmod +x /usr/local/bin/wait-for-it.sh

# Expose port (change if needed)
EXPOSE 8080

# Command to run the application
CMD ["/app/main"]

# ------------------------------
# Dev Stage (Chỉ dành cho môi trường phát triển)
# ------------------------------
FROM golang:1.23-alpine AS dev

# Thiết lập thư mục làm việc trong container
WORKDIR /app

# Cài đặt các công cụ cần thiết cho phát triển (git, bash, curl) và công cụ air để reload ứng dụng
RUN apk add --no-cache git bash curl \
    && go install github.com/air-verse/air@v1.61.7

# Copy file cấu hình air
COPY .air.toml ./

# Expose port 8080 cho việc phát triển
EXPOSE 8080

# Lệnh chạy cho môi trường phát triển
CMD ["air"]
