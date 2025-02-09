FROM golang:1.23.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .

RUN chmod +x /root/app

CMD ["./app"]