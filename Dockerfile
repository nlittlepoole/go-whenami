FROM golang:alpine


RUN cat /etc/apk/repositories
RUN sed -i -e 's/v3\.6/edge/g' /etc/apk/repositories
RUN echo 'http://nl.alpinelinux.org/alpine/edge/testing' >> /etc/apk/repositories
RUN cat /etc/apk/repositories
RUN apk update
RUN apk add --no-cache gcc musl-dev

RUN apk add sqlite-dev
RUN apk update && apk add sqlite

RUN apk add pkgconf
RUN apk add proj4-dev
RUN apk add zlib-dev
RUN apk add --update make
RUN apk add git
RUN apk add geos-dev
RUN apk add libxml2-dev
RUN apk add sudo

RUN wget http://www.gaia-gis.it/gaia-sins/freexl-1.0.5.tar.gz

RUN tar -xvzf freexl-1.0.5.tar.gz

RUN cd freexl-1.0.5 && ./configure && make && make install

RUN wget http://www.gaia-gis.it/gaia-sins/libspatialite-4.4.0-RC0.tar.gz

RUN tar -xvzf libspatialite-4.4.0-RC0.tar.gz

RUN cd libspatialite-4.4.0-RC0 && ./configure && make && sudo make install

RUN cp /usr/local/bin/* /usr/bin/
RUN cp -R /usr/local/lib/* /usr/lib/

ARG app_env
ENV APP_ENV $app_env

COPY ./ /go/src/github.com/nlittlepoole/whenami
WORKDIR /go/src/github.com/nlittlepoole/whenami/server

ENV PATH="$PATH:$GOROOT/bin:$GOPATH/bin"

RUN go get ./
RUN go build

CMD server

EXPOSE 80