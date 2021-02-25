FROM ubuntu:16.04



RUN apt update -y \
    && apt-get install iptables sudo -y


COPY demo /bin/demo


CMD ["demo"]