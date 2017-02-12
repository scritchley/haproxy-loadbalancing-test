FROM golang

MAINTAINER Simon Critchley

ADD . /go/src/code.simon-critchley.co.uk/server
WORKDIR /go/src/code.simon-critchley.co.uk/server

RUN go install code.simon-critchley.co.uk/server

ENTRYPOINT /go/bin/server

EXPOSE 5000
