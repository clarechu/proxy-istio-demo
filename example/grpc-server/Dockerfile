FROM ubuntu:16.04



RUN apt update -y \
    && apt-get install iptables wget telnet watch curl sudo -y


COPY grpc-server /bin/grpc-server


CMD ["grpc-server"]