FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY . .

ARG VERSION="1.0.0"

RUN go mod init myapp && go mod tidy && \
    go build -ldflags "-X main.version=$VERSION" -o myapp

FROM nginx:latest

COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

HEALTHCHECK --interval=30s --timeout=10s --retries=3 CMD curl -f http://localhost || exit 1

CMD ["nginx", "-g", "daemon off;"]
