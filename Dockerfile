FROM golang:alpine

WORKDIR /golang-docker

COPY ./src/. .

RUN go mod download

EXPOSE 9000
ENTRYPOINT go build  && ./golang-docker
