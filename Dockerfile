# This dockerfile uses extends image https://hub.docker.com/bar-counter/slog
# VERSION 1
# Author: sinlov
# dockerfile offical document https://docs.docker.com/engine/reference/builder/
# https://hub.docker.com/_/golang
FROM golang:1.17.13-buster as builder

ARG GO_PATH_SOURCE_DIR=/go/src/
WORKDIR ${GO_PATH_SOURCE_DIR}

RUN mkdir -p ${GO_PATH_SOURCE_DIR}/github.com/bar-counter/slog
COPY $PWD ${GO_PATH_SOURCE_DIR}/github.com/bar-counter/slog

RUN cd ${GO_PATH_SOURCE_DIR}/github.com/bar-counter/slog && \
    go mod download -x

RUN  cd ${GO_PATH_SOURCE_DIR}/github.com/bar-counter/slog && \
  CGO_ENABLED=0 \
  go build \
  -a \
  -installsuffix cgo \
  -ldflags '-w -s --extldflags "-static -fpic"' \
  -tags netgo \
  -o slog \
  main.go

# https://hub.docker.com/_/alpine
FROM alpine:3.17

ARG DOCKER_CLI_VERSION=${DOCKER_CLI_VERSION}

#RUN apk --no-cache add \
#  ca-certificates mailcap curl \
#  && rm -rf /var/cache/apk/* /tmp/*

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/github.com/bar-counter/slog/slog .
ENTRYPOINT ["/app/slog"]
# CMD ["/app/slog", "--help"]