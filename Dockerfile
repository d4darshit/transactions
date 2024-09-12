# Dockerfile for Go application
FROM golang:1.18-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules manifests
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app from the cmd folder
RUN go build -o transactions-api ./cmd/main.go

# Ensure the binary is executable
RUN chmod +x /app/transactions-api

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/app/transactions-api"]
