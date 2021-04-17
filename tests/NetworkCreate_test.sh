#!/bin/bash

. tests/const.sh

${DOCKERCTL}  --hostname=${HOST} network NetworkCreate --networkCreateBody.Name my-network

docker network rm my-network