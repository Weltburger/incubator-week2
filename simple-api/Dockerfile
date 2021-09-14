# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /usr/local/bin/simple-api ./cmd/main.go

EXPOSE 1323

CMD ["/usr/local/bin/simple-api"]
