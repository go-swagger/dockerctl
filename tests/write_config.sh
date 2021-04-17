#!/bin/bash

# used for manual testing to skip repetitive flags

# write config file to standard location
read -r -d '' CONTENT <<- EOM
hostname: "localhost:12345"
scheme: http
EOM

STANDARD_LOCATION="$HOME/.config/dockerctl"
mkdir -p $STANDARD_LOCATION
echo "$CONTENT" > "$STANDARD_LOCATION/config.yaml"