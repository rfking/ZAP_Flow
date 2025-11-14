FROM golang:1.24.3-alpine AS builder
WORKDIR /app

# Install gcc, build-base and SQLite dev libraries
RUN apk add --no-cache build-base sqlite-dev git

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies with improved retry logic and timeout settings
RUN go mod download -x 2>&1 || (sleep 3 && go mod download -x) || (sleep 5 && go mod download -x) || (sleep 10 && go mod download -x)

# Verify modules integrity
RUN go mod verify || true

# Copy all source code
COPY . .

# Build with CGO enabled for SQLite
RUN CGO_ENABLED=1 GOOS=linux CGO_LDFLAGS="-static" go build -a -installsuffix cgo -o whatsmiau main.go

# Final stage
FROM alpine:latest

RUN apk update && apk add --no-cache ca-certificates ffmpeg mailcap sqlite-libs

WORKDIR /app

# Create data directory with proper permissions
RUN mkdir -p /app/data && chmod 777 /app/data

# Copy binary from builder
COPY --from=builder /app/whatsmiau /app/whatsmiau

EXPOSE 8081

ENTRYPOINT ["/app/whatsmiau"]
