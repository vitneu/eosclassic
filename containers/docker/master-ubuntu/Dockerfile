FROM ubuntu:xenial

ENV PATH=/usr/lib/go-1.11/bin:$PATH

RUN \
  apt-get update && apt-get upgrade -q -y && \
  apt-get install -y --no-install-recommends golang-1.11 git make gcc libc-dev ca-certificates && \
  git clone --depth 1 --branch stable https://github.com/eosclassic/eosclassic && \
  (cd eosclassic && make eosc) && \
  cp eosclassic/build/bin/eosc /eosc && \
  apt-get remove -y golang-1.11 git make gcc libc-dev && apt autoremove -y && apt-get clean && \
  rm -rf /eosclassic

EXPOSE 8282
EXPOSE 25252

ENTRYPOINT ["/eosc"]
