FROM golang:alpine
MAINTAINER "XX"

ENV TERRAFORM_VERSION=0.8.1

RUN wget https://concourse.ci/downloads.html
RUN cat downloads.html

RUN apk add --update git bash

ENV TF_DEV=true

WORKDIR $GOPATH/src/github.com/hashicorp/terraform
RUN git clone https://github.com/mevansam/terraform.git ./ && \
    /bin/bash scripts/build.sh

WORKDIR $GOPATH
ENTRYPOINT ["terraform"]


