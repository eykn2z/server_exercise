FROM golang:latest
WORKDIR /go/src/work
COPY . /go/src/work
CMD go run server.go
EXPOSE 8080