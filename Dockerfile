FROM golang:1.6

RUN go get github.com/codegangsta/gin

ADD . /go/src/github.com/gotstago/todoapp-go-es
WORKDIR /go/src/github.com/gotstago/todoapp-go-es
RUN go get
