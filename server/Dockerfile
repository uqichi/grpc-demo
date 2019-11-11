FROM golang:1.13.3-alpine as builder

ENV GO111MODULE on
WORKDIR /go/src/demo
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /go/bin/demo-server -dflags="-s -w" .

FROM alpine

ENV TZ Asia/Tokyo
COPY --from=builder /go/bin/demo-server /demo-server
ENTRYPOINT ["/demo-server"]