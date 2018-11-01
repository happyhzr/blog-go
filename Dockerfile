FROM golang:1.11

WORKDIR /go/src/github.com/insisthzr/blog-go

COPY . .

RUN go build -o blog

ENTRYPOINT ./blog

EXPOSE 8080
