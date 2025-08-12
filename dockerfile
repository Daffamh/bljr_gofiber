FROM golang:alpine AS dev

RUN go install github.com/air-verse/air@latest

WORKDIR /app

EXPOSE 3000