FROM golang:1.6-onbuild
RUN echo "deb http://deb.debian.org/debian stretch main" >> /etc/apt/sources.list && \
    apt-get update && \
    apt-get install -y protobuf-compiler golang-goprotobuf-dev
