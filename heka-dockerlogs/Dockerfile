FROM debian
MAINTAINER William McGann <contact@WilliamMcGann.com>

ENV HEKA_TAG=v0.8.3 _VB=0_8_3

RUN apt-get update && apt-get install -y wget && rm -rf /var/lib/apt/lists/* && \
    wget --no-check-certificate https://github.com/mozilla-services/heka/releases/download/v0.8.3/heka-0_8_3-linux-amd64.tar.gz && \
    tar zxvf heka-${_VB}-linux-amd64.tar.gz && \
    rm heka-${_VB}-linux-amd64.tar.gz && \
    mv heka-${_VB}-linux-amd64 heka && \
    mkdir /conf

COPY conf/hekad.toml /conf/hekad.toml

CMD ["/heka/bin/hekad", "-config=/conf"]
