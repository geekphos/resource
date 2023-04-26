FROM golang:1.20 as builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim

RUN mkdir -p /app/configs
RUN mkdir -p /opt/assets

COPY --from=builder /src/_output/yoo /app
COPY --from=builder /src/configs/yoo.yaml /app/configs

WORKDIR /app

EXPOSE 8080

VOLUME ["/app/configs"]
VOLUME ["/opt/assets"]

CMD ["./yoo", "-c", "/app/configs/yoo.yaml"]