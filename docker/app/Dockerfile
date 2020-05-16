FROM nnvgm/grpc-common-example as builder

WORKDIR $GOPATH/src/github.com/nnvgm/grpc-client-example

COPY . .

RUN go get -d -v ./... && go build -o app

FROM alpine:3.10

COPY --from=builder /go/src/github.com/nnvgm/grpc-client-example/app /usr/bin/app

EXPOSE 8000

CMD ["app"]