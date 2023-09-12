FROM golang:latest as sdkcode

USER root

ENV LAATOO_SDK=/laatoo/sdk GO111MODULE=on

RUN mkdir -p $LAATOO_SDK

COPY src/ $LAATOO_SDK/

WORKDIR /laatoo/sdk/src

RUN go mod tidy 

RUN groupadd -r laatoo -g 1000

RUN useradd -u 1000 -ms /bin/bash -g laatoo -G sudo -p laatoo laatoo

RUN chown -R laatoo:laatoo /go $LAATOO_SDK

RUN tar -cf /tmp/gofolder.tar /go

RUN tar -cf /tmp/laatoo.tar /laatoo

FROM laatoodevelbase:latest

USER root

ENV LAATOO_SDK=/laatoo/sdk GO111MODULE=on

COPY --from=sdkcode /tmp/gofolder.tar /tmp/gofolder.tar

COPY --from=sdkcode /tmp/laatoo.tar /tmp/laatoo.tar

RUN tar -xf /tmp/gofolder.tar --directory /

RUN tar -xf /tmp/laatoo.tar --directory /

#RUN chown -R laatoo:laatoo /go $LAATOO_SDK

WORKDIR /go
