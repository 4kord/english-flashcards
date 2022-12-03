### STAGE 1: Build ###
FROM node:19.2 AS builder

WORKDIR /web

### SETTING ENV LIKE THIS BECAUSE VITE BUILDS THEM INTO CODE ###
RUN printf 'VITE_API_URL="http://localhost/api"\n' > .env

COPY web/package.json .
COPY web/vite.config.ts .
RUN npm install
COPY web/. .
RUN npm upgrade esbuild
RUN npm run build

### STAGE 2: Run ###
FROM nginx:1.23.2-alpine
COPY extras/nginx/default.conf /etc/nginx/conf.d
COPY extras/ssl /etc/nginx/ssl
RUN chmod -R 600 /etc/nginx/ssl
COPY --from=builder /web/dist /var/www/default
RUN chmod -R 755 /var/www/default