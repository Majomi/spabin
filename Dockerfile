FROM node:lts-alpine AS web-env
WORKDIR /build
COPY frontend/package*.json ./
RUN npm install
COPY frontend ./
RUN npm run build

FROM golang:1.13-alpine3.10 AS go-env
WORKDIR /build
COPY backend ./
COPY --from=web-env /build/dist ./static
RUN go generate && go build -o spabin -mod=vendor -v

FROM alpine:3.10
WORKDIR /app
COPY --from=go-env /build/spabin .
EXPOSE 3000
ENTRYPOINT [ "/app/spabin" ]