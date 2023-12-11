FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod tidy && go build cmd/server/server.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .
CMD [ "sh", "-c", "./server"]
