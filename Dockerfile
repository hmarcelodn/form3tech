FROM golang:1.16.7-alpine

RUN set -ex; \
    apk update; \
    apk add --no-cache git

WORKDIR /go/src/github.com/hmarcelodn/form3tech/

COPY . /go/src/github.com/hmarcelodn/form3tech/

CMD CGO_ENABLED=0 go test ./test/*.go -v -cover
