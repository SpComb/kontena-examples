FROM golang:1.7

ADD . /go/src/app

WORKDIR /go/src/app

RUN go install -v .

CMD ["/go/bin/app"]

EXPOSE 8000
