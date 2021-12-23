FROM  golang:1.16.0  as builder

WORKDIR /app
COPY . .
RUN GOOS=linux go build -o om cmd/web/main.go

FROM alpine:3
#FROM golang:1.15.0

WORKDIR /root
COPY --from=builder /app/om /root/app/om
COPY --from=builder /app/config/config.yaml /root/conf/default.config.yaml

ENV CONF_PATH=/root/conf/default.config.yaml

WORKDIR /root/app/

# https://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

CMD ["./om"]