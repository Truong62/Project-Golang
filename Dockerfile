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
# Dev Stage 
# ------------------------------
FROM golang:1.23-alpine AS dev

WORKDIR /app

RUN apk add --no-cache git bash curl \
    && go install github.com/air-verse/air@v1.61.7

COPY .air.toml ./

EXPOSE 8080

CMD ["air"]
