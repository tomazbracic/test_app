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


RUN git clone https://cda4669c2c40bcb2659fcc6b3644e3c52f9ea841@github.com/driveulu/kafka_dispatcher.git
RUN go get github.com/hashicorp/consul/api && \
    go get github.com/confluentinc/confluent-kafka-go/kafka

WORKDIR $GOPATH/src/github.com/driveulu/kafka_dispatcher

COPY send.go customer/send.go
RUN go install main.go

WORKDIR $GOPATH/bin

CMD main