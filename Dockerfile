# syntax=docker/dockerfile:1

FROM golang:1.16-alpine
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY *.go ./

RUN go build -o /cb-sample-server
CMD [ "/cb-sample-server" ]
# Build Server

# FROM golang:latest as builder
# ENV GOBIN /go/bin
# WORKDIR /go/src/github.com/QualiArts/cb-sample-server
# COPY / .
# RUN go test
# RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build

# # COPY theserver file to image from builder
# FROM alpine:latest
# WORKDIR /usr/local/bin/
# COPY --from=builder /go/src/github.com/QualiArts/cb-sample-server/sample .

# EXPOSE 8080

# CMD ["/usr/local/bin/cb-sample-server"]