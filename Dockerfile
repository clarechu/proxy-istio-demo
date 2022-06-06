FROM golang:1.16.6 as builder


ENV APP_BIN=/bin \
    PATH=${APP_BIN}:$PATH \
    TZ='Asia/Shanghai' \
    HOME=/root \
    GOPROXY=https://goproxy.io,direct


WORKDIR $HOME
COPY ./ $HOME

RUN CGO_ENABLED=0 GOOS=linux go build -o go-httpproxy .


FROM alpine:3.7 as prod

EXPOSE 8080

ENV APP_BIN=/bin \
    PATH=${APP_BIN}:$PATH \
    TZ='Asia/Shanghai' \
    HOME=/root

RUN  mkdir -p ${APP_BIN} ${APP_ROOT} \
     && sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
     && apk update \
     && apk upgrade \
     && apk --no-cache add ca-certificates iputils\
     && apk add -U tzdata ttf-dejavu busybox-extras curl bash git\
     && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

WORKDIR ${HOME}

COPY --from=builder $HOME/go-httpproxy ${APP_BIN}

CMD ["go-httpproxy"]