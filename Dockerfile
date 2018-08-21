FROM golang:1.10.3

RUN mkdir -p /go/src/github.com/user/go-socket
COPY . /go/src/github.com/user/go-socket
WORKDIR /go/src/github.com/user/go-socket

RUN go get ./
RUN go build

CMD socket

EXPOSE 5000
