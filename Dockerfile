FROM golang:latest
LABEL maintainer "Name <email>"
RUN go get github.com/Fullscript/go-web-server.git
ADD . /go/src/github.com/"github username"/myproject
RUN go install github.com/"github username"/myproject
ENTRYPOINT /go/bin/myproject
EXPOSE 8080
