# Use the official Golang image as a base image
FROM golang:1.23.0-alpine AS builder
## User service is as main service/gateway to connect with all other microservices in this project
## command i use to build is this
## run this on the internal directory
## docker build -f user/Dockerfile --build-arg PORT=3000 -t remahanrembulan/book-management-user:0.0.1 .

# Set the Current Working Directory inside the container
WORKDIR /app

# copy pkg folder
COPY pkg ../pkg

# Copy go mod and sum files
COPY user/go.mod user/go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY user .
# Use dumb init to grab pid 1 before fiber does
RUN apk add --no-cache dumb-init
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
COPY --from=builder ["/usr/bin/dumb-init", "/usr/bin/dumb-init"]
#copy the config file
# COPY --from=builder /app/.env .

# Expose port from  variable
ARG PORT
EXPOSE ${PORT}

#Use entry point to avoid fiber grab pid 1
ENTRYPOINT ["/usr/bin/dumb-init", "--"]

# Command to run the executable
CMD ["./main"]