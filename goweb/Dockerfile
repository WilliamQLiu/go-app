# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.8

ADD . /go/src/github.com/williamqliu/go-app/goweb
RUN mkdir /go-app
WORKDIR /go-app
COPY ./vendor/github.com/williamqliu/go-app/goweb/view /go-app/view
RUN go build -o main github.com/williamqliu/go-app/goweb 
ENTRYPOINT ["/go-app/main"]

# Document that the service listens on port 8080.
EXPOSE 8080
