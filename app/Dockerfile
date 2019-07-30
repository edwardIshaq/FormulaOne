FROM golang:1.9 as builder

COPY * /go/src/github.com/edwardIshaq/goVendorTest/app/
WORKDIR /go/src/github.com/edwardIshaq/goVendorTest/app

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /go/src/github.com/edwardIshaq/goVendorTest/app/app .
CMD ["./app"]
