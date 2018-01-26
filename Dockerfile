FROM golang

ADD . /go/src/github.com/insisthzr/blog-back

WORKDIR /go/src/github.com/insisthzr/blog-back 

RUN go build

ENTRYPOINT /go/src/github.com/insisthzr/blog-back/blog-back

EXPOSE 8080
