FROM golang:1.11

WORKDIR /go/src/app

COPY . .

RUN go build -o blog

ENTRYPOINT ./blog

EXPOSE 8080
