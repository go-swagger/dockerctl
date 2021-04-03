#!/bin/sh

. tests/const.sh

${DOCKERCTL} --hostname=${HOST} image ImageList || exit 1
