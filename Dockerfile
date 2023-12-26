
# Start from the official Golang image to build the binary.
FROM golang:alpine as builder

# Enable Go modules.
ENV GO111MODULE=on

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /app

# Copy go mod and sum files to download dependencies.
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the source from the current directory to the working directory inside the container.
COPY . .

# Build the Go app.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Start a new stage from scratch for a smaller final image.
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the previous stage.
COPY --from=builder /app/main .

# Command to run the executable.
CMD ["./main"]

# Expose port 8080 to the outside world.
EXPOSE 8080
