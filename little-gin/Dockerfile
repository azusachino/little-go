FROM golang:1.22 as builder

WORKDIR /app

COPY ./ /app

ENV GOPROXY https://goproxy.cn,direct

RUN ["go", "build"]

# runtime
FROM debian:buster

WORKDIR /app

COPY builder:/app/little-gin /app/app

ENTRYPOINT ["/app/app"]
