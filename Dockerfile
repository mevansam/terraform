FROM golang:alpine
MAINTAINER "XX"

ENV TERRAFORM_VERSION=0.8.1


RUN wget https://concourse.ci/downloads.html
RUN cat downloads.html

RUN http_proxy=http://172.24.248.61:8080 https_proxy=http://172.24.248.61:8080 apk add --update git bash

ENV TF_DEV=true

WORKDIR $GOPATH/src/github.com/hashicorp/terraform
RUN git clone https://github.com/mevansam/terraform.git ./ && \
    /bin/bash scripts/build.sh

WORKDIR $GOPATH
ENTRYPOINT ["terraform"]


