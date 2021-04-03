# TODO: this is not working yet

# ${DOCKERCTL} --hostname=${HOST} --debug image ImageCreate --fromImage alpine --tag 3.10.7

# curl -v --unix-socket /var/run/docker.sock \
#   -X POST "http://localhost/v1.41/images/create?fromImage=alpine&tag=3.10.7"