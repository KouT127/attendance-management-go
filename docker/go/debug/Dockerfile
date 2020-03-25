FROM golang:1.13-alpine as builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/
COPY .. ./

ENV GO111MODULE=off

RUN set -eux && \
  apk update && \
  apk add --no-cache git && \
  go get github.com/oxequa/realize && \
  go get -u github.com/go-delve/delve/cmd/dlv && \
  go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv

ENV GO111MODULE on

RUN go mod download

RUN go build -o attendance-management

#FROM alpine:3.11.3
#RUN apk add tzdata
#
#COPY --from=builder /go/src/attendance-management /go/src/attendance-management

RUN set -x && \
  addgroup go && \
  adduser -D -G go go && \
  chown -R go:go /go/src/attendance-management

CMD ["/go/src/attendance-management"]