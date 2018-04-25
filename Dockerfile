FROM golang

WORKDIR /go/src/github.com/insisthzr/blog-back 

COPY . .

RUN go build -o blog

ENTRYPOINT ./blog

EXPOSE 8080
