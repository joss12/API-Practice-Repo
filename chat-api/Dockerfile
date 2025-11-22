
FROM golang:1.24.0-alpine AS builder
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /bin/chat ./cmd/server

FROM gcr.io/distroless/base-debian12
COPY --from=builder /bin/chat /chat
ENV PORT=8082
EXPOSE 8082
ENTRYPOINT ["/chat"]

