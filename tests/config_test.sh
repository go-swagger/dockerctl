#!/bin/bash

. tests/const.sh

# test using config file

# write config file
read -r -d '' CONTENT <<- EOM
hostname: "localhost:12345"
scheme: http
EOM

# using custom file location
echo "$CONTENT" > temp.yaml
${DOCKERCTL} --config=temp.yaml --debug image ImageList || exit 1
rm -f temp.yaml

# using standard location and do not use --config flag
STANDARD_LOCATION="$HOME/.config/dockerctl"
mkdir -p $STANDARD_LOCATION
echo "$CONTENT" > "$STANDARD_LOCATION/config.yaml"
${DOCKERCTL} --debug image ImageList || exit 1
rm -rf $STANDARD_LOCATION
