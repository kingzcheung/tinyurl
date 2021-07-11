FROM golang:1.16-alpine3.12 as builder
ARG BUILD_MODE=build
RUN set -ex \
    && echo "https://mirrors.aliyun.com/alpine/v3.12/main/" > /etc/apk/repositories \
    && echo "https://mirrors.aliyun.com/alpine/v3.12/community/" >> /etc/apk/repositories \
#    && apk update --repository http://mirrors.aliyun.com/alpine/v3.12/main \
    && ls -lah /root \
    && apk add --no-cache --repository http://mirrors.aliyun.com/alpine/v3.12/main make yarn protobuf git \
    && GOPROXY=https://goproxy.cn go get google.golang.org/protobuf/cmd/protoc-gen-go \
                google.golang.org/grpc/cmd/protoc-gen-go-grpc

#
WORKDIR /go/src
ADD . .
RUN make ${BUILD_MODE}

FROM alpine:3.12 as prod
RUN apk add --no-cache --repository http://mirrors.aliyun.com/alpine/v3.12/main \
		ca-certificates \
		&& mkdir /data

WORKDIR /go/bin
COPY --from=builder /go/src/dist/bin/tinyurl .
CMD ["/go/bin/tinyurl"]