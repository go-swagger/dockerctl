#/bin/bash

. tests/const.sh

${DOCKERCTL} --hostname=${HOST} container ContainerList --all
