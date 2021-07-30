FROM golang:1.16.6-alpine

RUN apk update && apk add git
RUN mkdir /go/src/server

COPY src/ /go/src/server

WORKDIR /go/src/server
