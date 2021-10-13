#!/bin/bash
#
# Install solc 
#

set -e
set -u

if [ ! -e geth-versions/geth-v1.10.8/go-ethereum-v1.10.8/build/bin/geth ] ; then
    mkdir -p geth-versions/geth-v1.10.8
    cd geth-versions/geth-v1.10.8
    wget -O geth.tar.gz "https://github.com/ethereum/go-ethereum/archive/v1.10.8.tar.gz"
    tar -zxvf geth.tar.gz
    cd go-ethereum-v1.10.8
    PATH=/usr/lib/go-1.7/bin:$PATH make geth
    echo "Geth installed at $TRAVIS_BUILD_DIR/geth-versions/geth-v1.10.8/go-ethereum-v1.10.8/build/bin/geth"
else
    echo "Geth found at $TRAVIS_BUILD_DIR/geth-versions/geth-v1.10.8/go-ethereum-v1.10.8/build/bin/geth"
fi