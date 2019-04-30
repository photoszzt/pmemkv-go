#!/bin/bash
set -e
echo pass | sudo -S apt update -qq
echo pass | sudo -S apt install -y -qq libtbb-dev
echo pass | sudo -S add-apt-repository -y ppa:longsleep/golang-backports
echo pass | sudo -S apt update -qq
echo pass | sudo -S apt install -y -qq golang-go
git clone https://github.com/pmem/pmemkv pmemkv_src
cd pmemkv_src
mkdir bin
cd bin
cmake .. -DCMAKE_BUILD_TYPE=Release \
    -DCMAKE_INSTALL_PREFIX=/usr/local
make
echo pass | sudo -S cp libpmemkv.so /usr/local/lib/
echo pass | sudo -S cp ../src/pmemkv.h /usr/local/include/libpmemkv.h
cd /pmemkv-go
cd pmemkv_sys
go build
cd ../pmemkv
go build
cd ../example
go build
