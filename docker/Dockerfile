FROM debian:jessie

MAINTAINER cd <cd@atomist.com>
COPY ./gokit-service-demo ./
WORKDIR /opt/gokit

ENV DUMB_INIT_VERSION=1.2.0

RUN curl -s -L -O https://github.com/Yelp/dumb-init/releases/download/v$DUMB_INIT_VERSION/dumb-init_${DUMB_INIT_VERSION}_amd64.deb \
    && dpkg -i dumb-init_${DUMB_INIT_VERSION}_amd64.deb \
    && rm -f dumb-init_${DUMB_INIT_VERSION}_amd64.deb

EXPOSE 8080
ENTRYPOINT ["dumb-init", "./gokit-service-demo"]