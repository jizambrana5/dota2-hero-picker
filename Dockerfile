# Use an official Golang runtime as a parent image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY go.mod go.sum ./

# Install Go dependencies
RUN go mod download

# Copy the rest of your application code
COPY . .

COPY local.yml /app/internal/pkg

# Navigate to the directory containing the main file
WORKDIR /app/internal/pkg
# Build the Go application
RUN go build -o /app/myapp

# Expose a port for the application to listen on
EXPOSE 8081

# Run the application
CMD ["/app/myapp"]
