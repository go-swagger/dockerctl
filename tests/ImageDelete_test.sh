#!/bin/sh

. tests/const.sh

# delete a non existing image
test1(){
    echo "Delete unexisting image"
    ${DOCKERCTL} --hostname=${HOST} image ImageDelete --name alpine || return 1
}

# pull then delete
test2(){
    echo "Pull then delete"
    docker image pull alpine:3.10.7
    ${DOCKERCTL} --hostname=${HOST} image ImageDelete --name alpine:3.10.7 || return 1
}

test1 || exit 1
test2 || exit 1
