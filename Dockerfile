FROM golang

MAINTAINER Razil "zc6496359"

WORKDIR $GOPATH/src/sparta

ADD . $GOPATH/src/sparta

RUN go get -u github.com/go-sql-driver/mysql

RUN go build main.go

EXPOSE 5005

ENTRYPOINT ["./main"]