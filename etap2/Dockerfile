# syntax=docker/dockerfile:1.4

# Etap 1: Klonowanie repozytorium z GitHuba
FROM alpine AS gitclone

RUN apk add --no-cache git openssh

# Konfiguracja SSH na Windowsie
RUN mkdir -p /root/.ssh && chmod 0700 /root/.ssh
RUN touch /root/.ssh/known_hosts && ssh-keyscan github.com >> /root/.ssh/known_hosts

# Skorzystaj z przekazanego klucza SSH i sklonuj repo
RUN --mount=type=ssh git clone git@github.com:Neskka/PAwChO_lab5.git /app

# Etap 2: Serwer Nginx z aplikacją
FROM nginx:alpine
COPY --from=gitclone /app/etap2/dist /usr/share/nginx/html
