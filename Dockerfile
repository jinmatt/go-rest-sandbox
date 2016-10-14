FROM golang:alpine

COPY . /go/src/go-rest-sandbox

RUN go install go-rest-sandbox

ENTRYPOINT ["/go/bin/go-rest-sandbox"]

EXPOSE 8000
