FROM ubuntu:16.04




RUN apt update -y \
    && apt-get install iptables wget telnet watch curl sudo -y



COPY grpc-client /bin/grpc-client


CMD ["grpc-client"]