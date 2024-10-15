# Use the official Golang image as a base image
FROM golang:1.23.0-alpine AS builder

## command i use to build is this
## run this on the internal directory
## docker build -f category/Dockerfile --build-arg PORT=4041 -t remahanrembulan/book-management-category:0.0.1 .

# Set the Current Working Directory inside the container
WORKDIR /app

# copy pkg folder
COPY pkg ../pkg

# Copy go mod and sum files
COPY category/go.mod category/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY category .

# Build the Go app
# Use CGO_ENABLED=0 for a static binary and add -ldflags for stripping
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/server

# Start a new stage from scratch
FROM alpine:latest

# # add ca-cetificates
# RUN apk add --no-cache ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
# # copy the config file
# COPY --from=builder /app/.env .

# Expose port from  variable
ARG PORT
EXPOSE ${PORT}

# Command to run the executable
CMD ["./main"]