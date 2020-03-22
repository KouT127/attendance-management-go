FROM golang:1.13 as build

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/attendance-management/
COPY ./ /go/src/attendance-management/
COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build -o /bin/attendance-management

FROM alpine:3.11.3
RUN apk add tzdata
## RUN apk add --no-cache ca-certificates

COPY --from=build /bin/attendance-management /bin/attendance-management

CMD ["/bin/attendance-management"]