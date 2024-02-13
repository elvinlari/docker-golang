#
# Go
#
FROM golang:1.22.0 as goprog

ENV GO111MODULE=on

ADD . /go/src/github.com/elvinlari/docker-golang
WORKDIR /go/src/github.com/elvinlari/docker-golang

# RUN go mod init github.com/elvinlari/docker-golang

# RUN go get github.com/pkg/errors
# RUN go get github.com/gin-gonic/gin
# RUN go get github.com/sirupsen/logrus
# RUN go get github.com/coreos/go-oidc
# RUN go get gorm.io/driver/postgres
# RUN go get gorm.io/gorm

RUN apt-get update && \
    apt-get install -y vim && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN go install cmd/main.go

# RUN swag init -d "./" -g "cmd/main.go"


#
# Main App
#
FROM ubuntu:22.04

RUN apt-get update && \
    apt-get install -y supervisor vim && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Copy from goprog builds
RUN mkdir -p /usr/go/bin 
COPY --from=goprog /go/bin/ /usr/go/bin/

RUN mkdir -p /etc/scripts
ADD ./../scripts/supervisor.conf /etc/supervisor.conf
ADD ./../scripts/wait-for-it.sh /etc/scripts
ADD ./../scripts/go-app-entrypoint.sh /etc/scripts

RUN chmod +x /etc/scripts/wait-for-it.sh 
RUN chmod +x /etc/scripts/go-app-entrypoint.sh

ENTRYPOINT [ "/etc/scripts/go-app-entrypoint.sh" ]
EXPOSE 8080
