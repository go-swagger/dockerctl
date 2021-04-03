#!/bin/sh

# define constants used for tests

if [ "${CI}" != "1" ]; then
set -x
fi

# cmd exe
DOCKERCTL=cmd/cli/dockerctl

# host of docker engine. configured by socat
HOST=localhost:12345