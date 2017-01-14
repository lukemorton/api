FROM golang

LABEL name=api

ADD . /go/src/github.com/lukemorton/api

RUN go install github.com/lukemorton/api

CMD /go/bin/api

EXPOSE 3000
