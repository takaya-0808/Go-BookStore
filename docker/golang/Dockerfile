FROM golang:1.15.2-alpine

RUN apk update && apk add git

RUN mkdir /go/src/app

WORKDIR /go/src/app

ADD /src /go/src/app

RUN go get github.com/gin-gonic/gin

CMD [ "go", "run", "main.go" ]