FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Собираем бинарник, точка входа в backend/cmd/app/main.go
RUN go build -o backend ./cmd/app

EXPOSE 8080

CMD ["./backend"]