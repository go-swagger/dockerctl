# dockerfile for development with on unstable go-swagger
FROM docker:20.10.5-dind

# ENV GO111MODULE on

# Use unstable go-swagger build
# RUN go get github.com/youyuanwu/go-swagger@a6f47452c70598306bd7778e23c0df7114fb7b98
# RUN go-swagger version

RUN apk add curl socat make jq bash bash-completion

# copy completion script
COPY cmd/completion/dockerctl.bash-completion.sh /etc/bash_completion.d/dockerctl

# copy dockerctl exe
COPY cmd/dockerctl/dockerctl /usr/local/bin/dockerctl

RUN echo "#!/bin/bash" >> /root/.bashrc && \
    echo "source /etc/profile.d/bash_completion.sh " >> /root/.bashrc && \
    echo "source /etc/bash_completion.d/dockerctl" >> /root/.bashrc