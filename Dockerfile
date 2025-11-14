# Fase de build
FROM golang:1.24.3-alpine AS builder
WORKDIR /app

# Dependências C para SQLite
RUN apk add --no-cache build-base sqlite-dev git

# Go modules
COPY go.mod go.sum ./
RUN go mod download -x

# Copia código
COPY . .

# Build debug para facilitar erro detalhado
RUN CGO_ENABLED=1 go build -v -x -o whatsmiau main.go

# Fase final
FROM alpine:latest
RUN apk update && apk add --no-cache ca-certificates ffmpeg mailcap sqlite-libs
WORKDIR /app

RUN mkdir -p /app/data && chmod 777 /app/data

COPY --from=builder /app/whatsmiau /app/whatsmiau

EXPOSE 8081
ENTRYPOINT ["/app/whatsmiau"]
