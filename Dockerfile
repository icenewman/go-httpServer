FROM alpine:3.8

MAINTAINER dengxiaohui<dengxiaohui@zhiping.tech>

RUN echo "http://mirrors.aliyun.com/alpine/v3.8/main/" > /etc/apk/repositories
# 时区设置
RUN apk add --no-cache tzdata
# RUN apk add --no-cache git
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" >> /etc/timezone

RUN mkdir /app

WORKDIR /app

COPY main /app

#COPY http_server /app
#
#RUN go get github.com/gorilla/mux
#RUN go get github.com/hashicorp/consul/api

# RUN go build -i ./main.go

RUN chmod 777 ./main

EXPOSE 8080/tcp

ENTRYPOINT ["./main"]