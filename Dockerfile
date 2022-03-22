FROM golang:1.16.7-alpine

ENV ROOT=/app
WORKDIR ${ROOT}

RUN apk update && apk add --update git gcc build-base flex bison binutils-gold curl g++ gnupg libgcc linux-headers make && apk add --upgrade gmp-dev
ENV PBC_VERSION 0.5.14
ENV CGO_ENABLED=1

RUN set -ex \
    \
    && wget -O pbc.tar.gz "https://crypto.stanford.edu/pbc/files/pbc-$PBC_VERSION.tar.gz" \
    && mkdir -p /usr/src/pbc \
    && tar -xzC /usr/src/pbc --strip-components=1 -f pbc.tar.gz \
    && rm pbc.tar.gz \
    \
    && cd /usr/src/pbc \
    && ./configure \
    && make \
    && make install \
    && rm -rf /usr/src/pbc
COPY ./go.mod ./go.sum ./
RUN go mod download \
    # TODO: これらのモジュールは開発用なので本番ではインストールしない
    && go get -u github.com/cosmtrek/air
EXPOSE 4000