FROM golang:1.23-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -tags mysql -o /app/main ./cmd/ordersystem

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/main .

RUN chmod +x /app/main

COPY .env .
COPY configs.json .

EXPOSE 8000
EXPOSE 50051
EXPOSE 8080

CMD ["./main"]
