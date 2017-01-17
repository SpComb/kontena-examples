FROM golang:1.7

ADD . /go/src/app
RUN go install -v app

CMD ["/go/bin/app"]

EXPOSE 8000
