FROM golang:latest
RUN mkdir /go/src/charts_server
WORKDIR /go/src/charts_server/src/goTest/src
ADD . /go/src/charts_server

