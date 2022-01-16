FROM golang:1.17.2-alpine AS builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR $GOROOT/src/mod10homework

RUN go get github.com/prometheus/client_golang/prometheus
RUN go get github.com/prometheus/client_golang/prometheus/promhttp
COPY ./src .
RUN go build -o httpserver .

FROM scratch

COPY --from=builder $GOROOT/src/httpserver /

EXPOSE 8080

ENTRYPOINT ["/httpserver"]
