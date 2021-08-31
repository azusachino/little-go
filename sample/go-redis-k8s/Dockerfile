FROM golang:latest as builder
LABEL maintainer="az <azusachino@yahoo.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
COPY /go-redis-k8s .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

### New Stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
