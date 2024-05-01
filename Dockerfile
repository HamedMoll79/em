# Use an official Golang runtime as a parent image
FROM golang:1.20.4-alpine3.18

# Set the working directory
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Download dependencies
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download -x


# Build the binary executable
RUN go build -o main .

# Expose port 8080 for the application to listen on
EXPOSE 8088

# Start the application
CMD ["/app/main", "--migrate=up"]
