
FROM golang:1.26.3-alpine3.23 AS builder
WORKDIR /app

COPY go.mod go.sum
RUN go mod tidy

COPY . .
RUN go build -o main cmd/api/main.go

FROM alpine:latest AS runner
WORKDIR /root

COPY --from=builder /app/main .

COPY assets/mcc_risk.json app/assets/
COPY assets/normalization.json app/assets/
COPY .env app/

EXPOSE 9999

CMD ["./main"]