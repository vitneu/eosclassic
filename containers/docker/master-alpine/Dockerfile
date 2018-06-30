FROM alpine:3.7

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 --branch release/1.8 https://github.com/eosclassic/go-eosclassic && \
  (cd go-eosclassic && make geth) && \
  cp go-eosclassic/build/bin/geth /geth && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /go-eosclassic && rm -rf /var/cache/apk/*

EXPOSE 8282
EXPOSE 25252

ENTRYPOINT ["/geth"]
