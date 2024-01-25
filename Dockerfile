# Use the official Golang image as the base image
FROM golang:1.21.5-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod ./

# Download and install Go module dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the Go application
RUN go build -o main ./cmd

# Expose the port the application will run on
EXPOSE 8080

# Set environment variables
ENV PORT=8080

# Command to run the executable
CMD ["./main"]