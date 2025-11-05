FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.su[m] ./

RUN go mod download

COPY . .

# Build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o auth-system ./main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/auth-system .

EXPOSE 3003

CMD ["./auth-system"]
