# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./api ./api

RUN go build -ldflags "-s" -o ./bin/api ./api/cmd

## Deploy
FROM gcr.io/distroless/static-debian11

WORKDIR /

COPY --from=builder /app/bin/api /

EXPOSE 4000

USER nonroot:nonroot

CMD ["/api"]
