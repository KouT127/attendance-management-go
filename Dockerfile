FROM golang:1.13 as build

WORKDIR /go/src/attendance-management/
COPY ./ /go/src/attendance-management/

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /bin/attendance-management

FROM alpine:3.11.3
RUN apk add tzdata
## RUN apk add --no-cache ca-certificates

COPY --from=build /bin/attendance-management /bin/attendance-management

CMD ["/bin/attendance-management"]