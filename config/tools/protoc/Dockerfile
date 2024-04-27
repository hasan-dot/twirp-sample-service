FROM --platform=linux/amd64 golang:1.21

WORKDIR /go/src/twirp-service

# Squash a warning about world-writable dirs in PATH
RUN chmod o-w /go/bin /go/

RUN rm -rf /go/src/github.com/golang/protobuf
RUN go install github.com/golang/protobuf/protoc-gen-go@v1.5.2

RUN rm -rf /go/src/github.com/twitchtv/twirp
RUN go install github.com/twitchtv/twirp/protoc-gen-twirp@v8.1.3

RUN apt-get update && \
    apt-get -y install unzip
RUN curl -sL https://github.com/protocolbuffers/protobuf/releases/download/v21.9/protoc-21.9-linux-x86_64.zip -o /tmp/protoc.zip
RUN unzip /tmp/protoc.zip -d /tmp/protoc && \
  mv /tmp/protoc/bin/protoc /usr/local/bin/protoc && \
  mv /tmp/protoc/include/google /usr/local/include/

ADD config/tools/protoc/protoc /go/src/twirp-service/protoc

ENTRYPOINT ["/go/src/twirp-service/protoc"]
