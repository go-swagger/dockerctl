#/bin/sh

. tests/const.sh

docker pull hello-world:linux

${DOCKERCTL} --hostname=${HOST} --debug container ContainerCreate \
  --name my-hello --containerCreateBody.Domainname hello --containerCreateBody.Image hello-world:linux

