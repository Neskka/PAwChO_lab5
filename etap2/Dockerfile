FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .

ARG VERSION="1.0.0"

RUN go mod init myapp && go mod tidy && \
    go build -ldflags "-X main.version=$VERSION" -o myapp

FROM alpine:latest AS runtime

WORKDIR /root/
COPY --from=builder /app/myapp .
EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=10s --retries=3 CMD wget -qO- http://localhost:8080 || exit 1

CMD ["./myapp"]
