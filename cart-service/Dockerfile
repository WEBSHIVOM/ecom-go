# Use the official Go image as a base
FROM golang:1.20

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy Go Modules manifests
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Generate Swagger docs (this assumes `swag` is installed globally)
RUN swag init --parseDependency --parseInternal

# Build the Go app
RUN go build -o cart-service

# Expose port 8081
EXPOSE 8081

# Command to run the executable
CMD ["./cart-service"]
