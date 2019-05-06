# Copyright 2019 Intel Corporation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM ubuntu:bionic

RUN apt-get update \
  && apt-get install -y \
  wget \
  git \
  pkg-config \
  libssl-dev

RUN mkdir -p /go/
WORKDIR /go/

RUN wget https://dl.google.com/go/go1.12.2.linux-amd64.tar.gz \
  && tar -xf go1.12.2.linux-amd64.tar.gz \
  && mv ./go/* . \
  && rm go1.12.2.linux-amd64.tar.gz

RUN mkdir -p /golibraries

ENV GOROOT=/go
ENV GOBIN=/go/bin
ENV GOPATH=/golibraries
ENV PATH=$PATH:$GOBIN

RUN go get -u \
    github.com/jessevdk/go-flags \
    github.com/spacemonkeygo/openssl

WORKDIR /project/GoCryptoBenchmark/src

CMD go build -ldflags="-X main.DISTRIBUTION_NAME=$TOOL_NAME -X main.DISTRIBUTION_VERSION=$TOOL_VERSION" -o ../bin/$TOOL_NAME