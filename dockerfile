FROM golang:1.18.1-alpine3.15 AS builder
RUN apk add make
WORKDIR /code
COPY makefile makefile
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY cmd cmd
COPY internal internal
RUN make
