FROM golang:1.13.3-alpine as builder

RUN apk update && \
    apk add git gcc g++
ENV PKG_PATH /go/src/github.com/uqichi/grpc-demo
ENV APP_PATH $PKG_PATH/client
ENV GO111MODULE on
WORKDIR $PKG_PATH
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . $PKG_PATH
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -o /go/bin/app \
    -ldflags="-s -w -X main.revision=$(git rev-parse --short HEAD) -X main.version=$(git describe --tags)" $APP_PATH

FROM alpine

ENV TZ Asia/Tokyo
COPY --from=builder /go/bin/app /demo-client
ENTRYPOINT ["/demo-client"]