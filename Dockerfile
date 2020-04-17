FROM golang:1.13 as builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/
COPY ./ ./

RUN go mod download

RUN go build -o attendance-management ./server

FROM alpine:3.11.3
RUN apk add tzdata

COPY --from=builder /go/src/attendance-management /go/src/attendance-management

CMD ["/go/src/attendance-management"]