FROM alpine:latest

ENV GLIBC_VERSION 2.29-r0
ENV HFISH_VERSION 0.6.2
ENV HFISH_VERSION_NEW 0.6.3

# Download and install glibc
RUN apk update && \
  apk add --no-cache curl tzdata mysql-client && \
  cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
  curl -Lo /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub && \
  curl -Lo glibc.apk "https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-${GLIBC_VERSION}.apk" && \
  curl -Lo glibc-bin.apk "https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-bin-${GLIBC_VERSION}.apk" && \
  apk add glibc-bin.apk glibc.apk && \
  rm -rf glibc.apk glibc-bin.apk /var/cache/apk/*

RUN curl -Lo /tmp/HFish.tar.gz https://github.com/hacklcx/HFish/releases/download/${HFISH_VERSION}/HFish-${HFISH_VERSION}-linux-amd64.tar.gz && \
  tar -zxvf /tmp/HFish.tar.gz -C /tmp && \
  mv /tmp/HFish-${HFISH_VERSION}-linux-amd64 /tmp/HFish && \
  rm -f /tmp/HFish.tar.gz /tmp/HFish/LICENSE && \
  rm -f /tmp/HFish/HFish /tmp/HFish/static/data/js/data.js

RUN curl -Lo /tmp/HFish.tar.gz https://github.com/hacklcx/HFish/releases/download/${HFISH_VERSION_NEW}/HFish-${HFISH_VERSION_NEW}-linux-amd64.tar.gz && \
  tar -zxvf /tmp/HFish.tar.gz -C /tmp && \
  cp /tmp/HFish-${HFISH_VERSION_NEW}-linux-amd64/HFish /tmp/HFish/ && \
  cp /tmp/HFish-${HFISH_VERSION_NEW}-linux-amd64/static/data/js/data.js /tmp/HFish/static/data/js/ && \
  rm -rf /tmp/HFish-${HFISH_VERSION}-linux-amd64 && \
  rm -f /tmp/HFish.tar.gz

COPY Entrypoint.sh /Entrypoint.sh

RUN chmod +x /Entrypoint.sh

EXPOSE 21 22 23 69 3306 5900 6379 7879 8080 8081 8989 9000 9001 9200 11211

WORKDIR /opt

ENTRYPOINT ["/Entrypoint.sh"]
