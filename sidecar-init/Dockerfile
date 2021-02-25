FROM ubuntu:16.04


COPY istio-iptables /bin/istio-iptables

RUN apt update -y \
    && apt-get install iptables sudo -y \
    && chmod +x /bin/istio-iptables

CMD ["istio-iptables"]