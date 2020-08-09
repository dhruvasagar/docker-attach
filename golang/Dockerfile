FROM golang:latest AS builder

RUN go get github.com/go-delve/delve/cmd/dlv

WORKDIR /goapp
COPY . .
RUN go build -o goapp

FROM debian:buster

EXPOSE 1234

COPY --from=builder /go/bin/dlv .
COPY --from=builder /goapp/goapp .

CMD ["./dlv", "exec", "./goapp"]
