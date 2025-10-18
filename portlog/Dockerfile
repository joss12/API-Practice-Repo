# syntax=docker/dockerfile:1
FROM golang:1.24

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o portlog

EXPOSE 9000

CMD ["./portlog"]
