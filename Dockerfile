FROM golang
LABEL name="api"
ADD . /go/src/github.com/lukemorton/api
RUN go get github.com/gorilla/handlers
RUN go install github.com/lukemorton/api/server
CMD /go/bin/server
EXPOSE 3000
