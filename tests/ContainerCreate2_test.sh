#/bin/sh

. tests/const.sh

# specify complicated options and only dry run

${DOCKERCTL} --hostname=${HOST} --debug --dry-run container ContainerCreate \
  --name my-hello \
  --containerCreateBody.ArgsEscaped \
  --containerCreateBody.AttachStderr \
  --containerCreateBody.Domainname hello \
  --containerCreateBody.Healthcheck.Interval 10 \
  --containerCreateBody.Healthcheck.Retries 11 \
  --containerCreateBody.Healthcheck.StartPeriod 12 \
  --containerCreateBody.Healthcheck.Timeout 13 \
  --containerCreateBody.Image hello-world:linux


