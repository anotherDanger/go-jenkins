FROM golang:1.24.1-alpine AS builder

RUN mkdir app

WORKDIR /app

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]