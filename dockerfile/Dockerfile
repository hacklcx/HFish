FROM alpine:latest

ENV GLIBC_VERSION 2.29-r0
ENV HFISH_VERSION 0.2

# Download and install glibc
RUN apk update && \
  apk add --no-cache curl tzdata && \
  cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
  curl -Lo /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub && \
  curl -Lo glibc.apk "https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-${GLIBC_VERSION}.apk" && \
  curl -Lo glibc-bin.apk "https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-bin-${GLIBC_VERSION}.apk" && \
  apk add glibc-bin.apk glibc.apk && \
  rm -rf glibc.apk glibc-bin.apk /var/cache/apk/*

RUN curl -Lo /tmp/hfish.tar.gz https://github.com/hacklcx/HFish/releases/download/${HFISH_VERSION}/hfish-${HFISH_VERSION}-linux-amd64.tar.gz && \
  tar -zxvf /tmp/hfish.tar.gz -C /tmp && \
  mv /tmp/hfish-${HFISH_VERSION}-linux-amd64 /tmp/hfish && \
  rm -f /tmp/hfish.tar.gz /tmp/hfish/LICENSE /tmp/hfish/README.md

COPY Entrypoint.sh /Entrypoint.sh

RUN chmod +x /Entrypoint.sh

EXPOSE 21 22 23 3306 6379 7879 9000 9001 9002

ENTRYPOINT ["/Entrypoint.sh"]

