FROM golang
WORKDIR /go/src/httpserver
COPY . .
RUN go env -w GO111MODULE=auto
RUN go mod init
RUN go build -o /go/bin/httpserver main.go
WORKDIR ../../bin/
ENTRYPOINT ["./httpserver"]
ENV MY_SERVICE_PORT=8080
LABEL label="httpserver"
EXPOSE 8080


