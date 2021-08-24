FROM ubuntu:16.04

MAINTAINER suscan

ENV LC_ALL C.UTF-8
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' > /etc/timezone

RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list

RUN set -x \
   && apt-get clean \
   && apt-get update -y \
   && apt-get install -y nmap\
   && apt-get -qq update \
   && apt-get -qq install -y --no-install-recommends ca-certificates curl


WORKDIR /build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY . .

EXPOSE 18000

CMD ["./main"]
