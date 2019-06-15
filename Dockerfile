FROM golang

WORKDIR $GOPATH/src/sparta

ADD . $GOPATH/src/sparta

RUN env=production

RUN go get -d -v ./...

RUN go build main.go

EXPOSE 5005

ENTRYPOINT ["./main"]