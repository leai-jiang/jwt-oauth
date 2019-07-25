FROM golang

WORKDIR $GOPATH/src/sparta

ADD . $GOPATH/src/sparta

RUN export env=production
RUN export NODE_ENV=production

RUN go get -d -v ./...

RUN go build main.go

EXPOSE 5005

ENTRYPOINT ["./main"]
