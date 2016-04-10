FROM golang:1.6

RUN go get github.com/codegangsta/gin
RUN go get -u -v github.com/nsf/gocode
RUN go get -u -v github.com/rogpeppe/godef
RUN go get -u -v github.com/golang/lint/golint
RUN go get -u -v github.com/lukehoban/go-find-references
RUN go get -u -v github.com/lukehoban/go-outline
RUN go get -u -v sourcegraph.com/sqs/goreturns
RUN go get -u -v golang.org/x/tools/cmd/gorename
RUN go get -u -v github.com/tpng/gopkgs
RUN go get -u -v github.com/newhook/go-symbols

ADD . /go/src/github.com/gotstago/todoapp-go-es
WORKDIR /go/src/github.com/gotstago/todoapp-go-es
RUN go get
