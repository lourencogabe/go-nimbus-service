FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/api/main.go

# Imagem final leve baseada em Debian Slim (funciona direto)
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]