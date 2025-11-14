FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install gcc and SQLite dev libraries
RUN apk add --no-cache build-base sqlite-dev git

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies with retry logic
RUN go mod download || (sleep 2 && go mod download) || (sleep 5 && go mod download)

# Copy all source code
COPY . .

# Clean Go cache and rebuild modules
RUN go clean -modcache
RUN go mod download

# Build with CGO enabled for SQLite
RUN CGO_ENABLED=1 GOOS=linux CGO_LDFLAGS="-static" go build -a -installsuffix cgo -o whatsmiau main.go

# Final stage
FROM alpine:latest

RUN apk update && apk add --no-cache ca-certificates ffmpeg mailcap sqlite-libs

WORKDIR /app

# Create data directory
RUN mkdir -p /app/data && chmod 777 /app/data

# Copy binary from builder
COPY --from=builder /app/whatsmiau /app/whatsmiau

EXPOSE 8081

ENTRYPOINT ["/app/whatsmiau"]
