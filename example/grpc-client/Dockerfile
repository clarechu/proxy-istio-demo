FROM ubuntu:16.04



RUN apt update -y \
    && apt-get install iptables sudo -y


COPY grpc-client /bin/grpc-client


CMD ["grpc-client"]