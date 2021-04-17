#!/bin/bash

. tests/const.sh

docker pull hello-world:linux

${DOCKERCTL} --hostname=${HOST} distribution DistributionInspect --name hello-world:linux