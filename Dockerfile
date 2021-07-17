FROM golang:latest

RUN mkdir /app

WORKDIR /app

RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/gorm

COPY go.mod .
COPY go.sum .