# Use the official Golang image for building the Go application
FROM golang:1.20 AS builder

# Set the working directory
WORKDIR /app

# Copy the go.mod file to the container
COPY go.mod ./

# Download dependencies and generate go.sum if missing
RUN go mod tidy

# Copy the source code into the container
COPY main.go ./

# Build the Go application for Linux x86_64 with static linkage
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /request-debugger

# Use a minimal base image for the final container
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /request-debugger /request-debugger

# Expose port 80
EXPOSE 80

# Set the entrypoint
ENTRYPOINT ["/request-debugger"]

