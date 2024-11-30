FROM golang:1.23 as builder

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init

RUN go mod tidy

RUN go build -o nooter main.go

EXPOSE 8080

CMD ["./nooter"]
