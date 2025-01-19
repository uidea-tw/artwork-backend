# Stage 1: Build stage
FROM golang:1.23.4-alpine AS builder

# Install dependencies and tools
RUN apk add --no-cache git

# Create and set the working directory
WORKDIR /src

# Copy go.mod and go.sum, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire application source code
COPY . .

# Build the application
RUN go build -o /bin/server

# Stage 2: Runtime stage
FROM alpine:3.18

# Set up working directory and copy the built binary from builder
WORKDIR /app
COPY --from=builder /bin/server /bin/server

# coyp configs from builder
COPY --from=builder /src/configs /app/configs

# Expose the application port (adjust as needed)
EXPOSE 8080

# Set the default command
CMD ["/bin/server"]