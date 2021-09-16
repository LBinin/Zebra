FROM golang:alpine AS build-env

ADD . /tmp/shop
RUN set -ex \
    && cd /tmp/shop \
    && export GOPROXY="https://goproxy.io,direct" \
    && go build -o /usr/bin/shop  cmd/shop/main.go

FROM alpine
MAINTAINER kilingzhang <i@kilingzhang.com>

RUN set -ex \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
    && apk add --no-cache tzdata \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

COPY --from=build-env /usr/bin/shop /usr/bin/shop
COPY docker-entrypoint.sh .

ENTRYPOINT ["sh", "docker-entrypoint.sh"]