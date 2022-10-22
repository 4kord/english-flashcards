FROM node:18-alpine3.15 AS builder

WORKDIR /web
COPY /web .

RUN yarn install && yarn build

FROM nginx:1.23.1-alpine

WORKDIR /etc/nginx/conf.d
RUN rm -rf ./*

COPY /docker/nginx/conf.d/default.conf .

WORKDIR /etc/nginx/ssl
RUN rm -rf ./*
RUN chmod -R 600 /etc/nginx/ssl

COPY /docker/nginx/ssl .

WORKDIR /var/www/default
RUN rm -rf ./*
RUN chmod -R 755 /var/www/default

COPY --from=builder /web/build .
