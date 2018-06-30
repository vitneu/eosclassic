# Build Geth in a stock Go builder container
FROM golang:1.10-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /go-eosclassic
RUN cd /go-eosclassic && make geth

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-eosclassic/build/bin/geth /usr/local/bin/

EXPOSE 8282 8546 25252 25252/udp
ENTRYPOINT ["geth"]
