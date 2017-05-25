# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.8

WORKDIR home/will/go/src/github.com/williamqliu/go-app/goweb
RUN mkdir /myapp
COPY . /myapp

# Build the go-app command inside the container
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
ENTRYPOINT ["/myapp/goweb"]

# Document that the service listens on port 8080.
EXPOSE 8080

