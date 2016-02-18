FROM golang:latest

ENV GO15VENDOREXPERIMENT 1
EXPOSE  8080

RUN mkdir /go/src/dockito
WORKDIR /go/src/dockito
COPY ./ /go/src/dockito/

RUN go build .
CMD ./dockito
