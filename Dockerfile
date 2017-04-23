FROM golang:alpine

RUN mkdir -p /go/src/github.com/bturbes/demo_api
WORKDIR /go/src/github.com/bturbes/demo_api

COPY . /go/src/github.com/bturbes/demo_api
RUN go build

CMD [ "demo_api" ]