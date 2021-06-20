#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app

COPY go.* .
RUN go mod download

COPY . .
RUN go build -o /go/bin/app ./cmd/web/*.go

#final stage
FROM alpine:latest

LABEL Name=movies Version=0.0.1
ENV TIMEZONE Europe/Istanbul
EXPOSE 4000

# setup timezone
RUN apk update && apk upgrade && \
    apk add --no-cache ca-certificates && \
    apk add -U --no-cache --virtual .build-deps tzdata && \
    cp /usr/share/zoneinfo/$TIMEZONE /etc/localtime && echo $TIMEZONE > /etc/timezone && \
    apk del .build-deps && \
    rm -rf /var/cache/apk/*

COPY --from=builder /go/bin/app /app
COPY views /views
COPY static /static

ENTRYPOINT /app
