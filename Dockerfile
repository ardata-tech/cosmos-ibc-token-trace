# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder

# Install git for go mod
RUN apk add --no-cache git

WORKDIR /app

# Cache go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code and build
COPY . .
RUN go build -o /cosmos-ibc-check

# Stage 2: Create the final image
FROM alpine:latest

# Install ca-certificates
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /cosmos-ibc-check /cosmos-ibc-check

CMD ["/cosmos-ibc-check"]

