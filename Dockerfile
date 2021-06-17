FROM golang:latest

LABEL maintainer="Lado <ladorayhannajib@gmail.com>"

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY Makefile .

RUN go mod download

COPY . .

ENV PORT 8000