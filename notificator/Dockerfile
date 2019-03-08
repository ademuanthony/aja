FROM golang

RUN mkdir -p /go/src/github.com/ademuanthony/aja

ENV GO111MODULE=on

ADD . /go/src/github.com/ademuanthony/aja
WORKDIR /go/src/github.com/ademuanthony/aja/notificator

RUN go mod download

# RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/ademuanthony/aja/notificator/cmd  -watch github.com/ademuanthony/aja/notificator
