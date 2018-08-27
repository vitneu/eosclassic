# Build EOSC in a stock Go builder container
FROM golang:1.11-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /eosclassic
RUN cd /eosclassic && make eosc

# Pull EOSC into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /eosclassic/build/bin/eosc /usr/local/bin/

EXPOSE 8282 8285 25252 25252/udp
ENTRYPOINT ["eosc"]
