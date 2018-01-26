FROM golang

ADD . /go/src/github.com/insisthzr/blog-back

WORKDIR /go/src/github.com/insisthzr/blog-back 

RUN go build

ENTRYPOINT /go/bin/blog-back

EXPOSE 8080
