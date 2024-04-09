FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

# RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine:latest
ENV REDIRECT_URL="http://example.com"

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
