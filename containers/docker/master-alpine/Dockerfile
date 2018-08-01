FROM alpine:3.7

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 --branch release/1.8 https://github.com/eosclassic/eosclassic && \
  (cd eosclassic && make eosc) && \
  cp eosclassic/build/bin/eosc /eosc && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /eosclassic && rm -rf /var/cache/apk/*

EXPOSE 8282
EXPOSE 25252

ENTRYPOINT ["/eosc"]
