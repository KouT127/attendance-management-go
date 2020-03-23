FROM golang:1.13 as builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/attendance-management/
COPY ./ /go/src/attendance-management/
COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build -o app

FROM alpine:3.11.3
RUN apk add tzdata

COPY --from=builder /go/src/attendance-management/app /go/src/attendance-management/app
COPY --from=builder /go/src/attendance-management/configs/ /go/src/attendance-management/configs/

CMD ["/go/src/attendance-management/app"]