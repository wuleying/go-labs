#!/usr/bin/env bash

echo "===images==="
docker images

docker rmi hyperledger/fabric-tools \
           hyperledger/fabric-couchdb \
           hyperledger/fabric-kafka \
           hyperledger/fabric-orderer \
           hyperledger/fabric-peer \
           hyperledger/fabric-ccenv \
           hyperledger/fabric-ca \
           hyperledger/fabric-baseimage \
           hyperledger/fabric-baseos

docker rmi 0403fd1c72c7 \
           2fbdbf3ab945 \
           dbd3f94de4b5 \
           e317ca5638ba \
           6830dcd7b9b5 \
           7182c260a5ca \
           a15c59ecda5b \
           9f2e9ec7c527 \
           4b0cab202084

echo "===images==="
docker images