FROM golang:1.16 as builder

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /simple_http

COPY . .

RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -o simple_httpserver .

FROM alpine:latest as prod

WORKDIR /root/

COPY --from=builder /simple_http/simple_httpserver  .

EXPOSE 80

CMD ["./simple_httpserver"]
