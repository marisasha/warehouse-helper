FROM golang:1.25.6-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG SERVICE
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o app ./cmd/${SERVICE}/main.go


FROM alpine:3.20

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/app .
COPY --from=builder /app/configs ./configs
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/docs ./docs

EXPOSE 8000

CMD ["./app"]
