FROM alpine:latest

RUN mkdir -p /hello/go

FROM scratch

WORKDIR /hello/go

COPY hello .