# Use the official Golang image as the base image
FROM golang:latest as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and download dependencies
COPY go.mod .
# COPY go.sum .
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the built executable from the builder stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
