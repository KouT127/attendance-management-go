FROM golang:1.13 as builder

WORKDIR /go/src/github.com/KouT127/attendance-management/
COPY . .

#RUN go mod download

WORKDIR /go/src/github.com/KouT127/attendance-management/backend/

RUN CGO_ENABLED=0 GOOS=linux go build -v -o main

FROM alpine

COPY --from=builder /go/src/github.com/KouT127/attendance-management/backend/main /main

CMD ["/main"]