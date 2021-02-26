FROM openjdk:8-jre-alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories

RUN apk add --no-cache curl grep sed unzip bash curl vim wget openssh busybox-extras

WORKDIR /app

COPY target/demo-0.0.1-SNAPSHOT.jar /app/

CMD java -Xms256m -Xmx256m -jar demo-0.0.1-SNAPSHOT.jar