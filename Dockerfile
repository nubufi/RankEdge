# Use the official Golang image as the base image
FROM golang:1.23rc2-alpine3.19 AS builder
# Set the Current Working Directory inside the container

RUN go install github.com/air-verse/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .


# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["air"]

