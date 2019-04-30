#!/bin/bash
set -e

if [[ -z "$HOST_WORKDIR" ]]; then
    HOST_WORKDIR=$(readlink -f .)
fi

chmod -R a+w $HOST_WORKDIR
WORKDIR=/pmemkv-rs

docker run --privileged=true \
    -v $HOST_WORKDIR:$WORKDIR \
    -v /etc/localtime:/etc/localtime \
    -w $WORKDIR \
    -it pmem/pmemkv:ubuntu-18.04 \
    ./run_build.sh
