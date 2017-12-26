#!/usr/bin/env bash

# pull fabric images
ARCH=x86_64
BASEIMAGE_RELEASE=0.3.1
PROJECT_VERSION=1.0.0
IMG_TAG=1.0.0

echo "Downloading fabric images from DockerHub...with tag = ${IMG_TAG}... need a while"
# TODO: we may need some checking on pulling result?
docker pull hyperledger/fabric-peer:$ARCH-$IMG_TAG
docker pull hyperledger/fabric-orderer:$ARCH-$IMG_TAG
docker pull hyperledger/fabric-ca:$ARCH-$IMG_TAG
docker pull hyperledger/fabric-tools:$ARCH-$IMG_TAG
docker pull hyperledger/fabric-couchdb:$ARCH-$IMG_TAG
docker pull hyperledger/fabric-kafka:$ARCH-$IMG_TAG
docker pull hyperledger/fabric-ccenv:$ARCH-$PROJECT_VERSION
docker pull hyperledger/fabric-baseimage:$ARCH-$BASEIMAGE_RELEASE
docker pull hyperledger/fabric-baseos:$ARCH-$BASEIMAGE_RELEASE

# Only useful for debugging
# docker pull yeasy/hyperledger-fabric

echo "===Re-tagging images to *latest* tag"
docker tag hyperledger/fabric-peer:$ARCH-$IMG_TAG hyperledger/fabric-peer
docker tag hyperledger/fabric-orderer:$ARCH-$IMG_TAG hyperledger/fabric-orderer
docker tag hyperledger/fabric-ca:$ARCH-$IMG_TAG hyperledger/fabric-ca
docker tag hyperledger/fabric-tools:$ARCH-$IMG_TAG hyperledger/fabric-tools
docker tag hyperledger/fabric-couchdb:$ARCH-$IMG_TAG hyperledger/fabric-couchdb
docker tag hyperledger/fabric-kafka:$ARCH-$IMG_TAG hyperledger/fabric-kafka
docker tag hyperledger/fabric-ccenv:$ARCH-$PROJECT_VERSION hyperledger/fabric-ccenv
docker tag hyperledger/fabric-baseimage:$ARCH-$BASEIMAGE_RELEASE hyperledger/fabric-baseimage
docker tag hyperledger/fabric-baseos:$ARCH-$BASEIMAGE_RELEASE hyperledger/fabric-baseos

docker images