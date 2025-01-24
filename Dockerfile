# Use the official Golang image with Go 1.23 from Docker Hub
FROM golang:1.23-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the entire project into the container
COPY . .

# Set the working directory to where your main.go file is located
WORKDIR /app/cmd/server

# Build the Go app
RUN go build -o /app/news-api .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the Go app
CMD ["/app/news-api"]


