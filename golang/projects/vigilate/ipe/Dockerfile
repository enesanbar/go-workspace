FROM alpine:latest

ENV TIMEZONE Europe/Istanbul
EXPOSE 4001
ENV VERSION 1.2.0

# setup timezone
RUN apk update && apk upgrade && \
    apk add --no-cache ca-certificates && \
    apk add -U --no-cache --virtual .build-deps tzdata wget tar && \
    cp /usr/share/zoneinfo/$TIMEZONE /etc/localtime && echo $TIMEZONE > /etc/timezone && \
    wget https://github.com/dimiro1/ipe/releases/download/v${VERSION}/ipe_v${VERSION}_linux_amd64.tar.gz && \
    tar xzvf ipe_v${VERSION}_linux_amd64.tar.gz && \
    cp ipe /usr/local/bin/ipe && \
    chmod +x /usr/local/bin/ipe && \
    apk del .build-deps && \
    rm -rf /var/cache/apk/*

RUN mkdir /etc/ipe
COPY config.json /etc/ipe/config.json

ENTRYPOINT /usr/local/bin/ipe -config /etc/ipe/config.json

