# Use the official Go image as a base
FROM golang:1.20

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o payment-service

# Expose port 8082
EXPOSE 8082

# Command to run the executable
CMD ["./payment-service"]
