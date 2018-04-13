FROM golang:1.9.1-alpine
RUN apk update && apk upgrade && apk add --no-cache bash \
                                                    git \
                                                    alpine-sdk \
                                                    python \
                                                    python-dev \
                                                    py-pip \
                                                    build-base \
                                                    openssh
WORKDIR /

ENV GOBIN=$GOPATH/bin

WORKDIR $GOPATH/src/github.com/tomazbracic/


#
# FIX BELOW
#


RUN git clone https://github.com/tomazbracic/test_app.git
RUN go get github.com/gocql/gocql

WORKDIR $GOPATH/src/github.com/tomazbracic/test_app

RUN go install test_app.go

WORKDIR $GOPATH/bin

CMD test_app