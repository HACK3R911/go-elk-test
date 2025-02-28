FROM golang:1.24.0-alpine3.21

COPY go.mod go.sum /go/src/github.com/HACK3R911/go-elk-test/
WORKDIR /go/src/github.com/HACK3R911/go-elk-test
RUN go mod download
COPY . /go/src/github.com/HACK3R911/go-elk-test
RUN go build -o /usr/bin/go-elk-test github.com/HACK3R911/go-elk-test/cmd/api

EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/go-elk-test"]