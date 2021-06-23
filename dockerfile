FROM alpine:latest

COPY ./hfish /opt/hfish

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

EXPOSE 4433 4434

WORKDIR /opt/hfish

CMD ["./server"]
