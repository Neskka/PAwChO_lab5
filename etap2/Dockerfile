# syntax=docker/dockerfile:1.4

# Etap 1: Klonowanie repozytorium z GitHuba
FROM alpine AS gitclone

RUN apk add --no-cache git openssh

# Konfiguracja SSH na Windowsie
RUN mkdir -p /root/.ssh && chmod 0700 /root/.ssh
RUN touch /root/.ssh/known_hosts && ssh-keyscan github.com >> /root/.ssh/known_hosts

RUN --mount=type=ssh git clone git@github.com:Neskka/PAwChO_lab5.git /app

FROM node:alpine AS nodeapp

WORKDIR /app
COPY --from=gitclone /app/ /app
WORKDIR /app/etap2

RUN npm install

# Etap 2: Serwer Nginx z aplikacją
FROM nginx:alpine

COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

COPY --from=nodeapp /app/etap2 /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]