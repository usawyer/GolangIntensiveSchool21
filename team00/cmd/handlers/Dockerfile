FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY . .
RUN go mod tidy && go build cmd/handlers/handlers.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/handlers .
COPY web/ /app/web
CMD [ "sh", "-c", "./handlers"]
