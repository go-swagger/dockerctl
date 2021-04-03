#!/bin/sh

. tests/const.sh

# existing image
${DOCKERCTL} --hostname=${HOST} --debug image ImageHistory --name alpine:3.10.7

# non-existing image
${DOCKERCTL} --hostname=${HOST} --debug image ImageHistory --name alpine:not-exist