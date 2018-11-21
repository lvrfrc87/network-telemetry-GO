FROM golang:1.11.2-alpine3.8
WORKDIR /go/src/ping
COPY src/ping /go/src/ping
COPY ./src/ping/var/.credentials.json /root/
RUN apk update && \
    apk --no-cache add git=2.18.1-r0
RUN go get -d -v github.com/influxdata/influxdb/client/v2 && \
    go get -d -v gopkg.in/yaml.v2 && \
    go install -v parser/json_parser_alpine.go && \
    go install -v cmd/ping_cmd_alpine.go && \
    go install -v yaml/yaml.go && \
    go install -v credentials/credentials.go && \
    gofmt -w ../ping/ && \
    go build -v main.go
CMD ["./main"]
