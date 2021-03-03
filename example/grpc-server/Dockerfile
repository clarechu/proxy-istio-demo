FROM ubuntu:16.04



RUN apt update -y \
    && apt-get install iptables sudo -y


COPY grpc-server /bin/grpc-server


CMD ["grpc-server"]