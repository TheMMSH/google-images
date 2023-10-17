FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/conf/config.yaml conf/config.yaml

CMD ["./main"]

