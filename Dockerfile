# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:alpine

RUN apk add make
# Set the Current Working Directory inside the container
RUN mkdir -p /go/src/github.com/gfleury/squaas
WORKDIR /go/src/github.com/gfleury/squaas

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN make


FROM node:lts

RUN mkdir -p /app
WORKDIR /app

# Copy build go application
COPY --from=0 /go/src/github.com/gfleury/squaas .

# Build the NodeJS frontend app
RUN make frontend/build


FROM alpine:latest

RUN mkdir -p /app
RUN mkdir -p /app/frontend/build
RUN mkdir -p /app/config


WORKDIR /app

# Copy build go application
COPY --from=1 /app/squaas .
COPY --from=1 /app/frontend/build frontend/build

# Volumify config directory
VOLUME ["/app/config"]

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./squaas"]



