#!/bin/bash

. tests/const.sh

docker network create my-network

${DOCKERCTL}  --hostname=${HOST} network NetworkDelete --id my-network
