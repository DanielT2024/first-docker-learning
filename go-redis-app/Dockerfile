# Use a Go base image
FROM golang:1.18-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files (if you had a go.sum)
COPY go.mod ./
# COPY go.sum ./

# Download Go modules
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
# CGO_ENABLED=0 disables CGO, making the binary static and easier to containerize
RUN CGO_ENABLED=0 go build -o /go-redis-app

# Use a minimal base image for the final container
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the compiled application from the builder stage
COPY --from=builder /go-redis-app .

# Expose the port our Go app listens on
EXPOSE 8080

# Run the compiled application
CMD ["./go-redis-app"]