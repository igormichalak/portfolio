# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-alpine AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./api/ ./

RUN go build -ldflags "-s" -o ./bin/api ./cmd

## Deploy
FROM gcr.io/distroless/static-debian11

WORKDIR /

COPY --from=builder /app/bin/api /

EXPOSE 4000

USER nonroot:nonroot

CMD ["/api"]
