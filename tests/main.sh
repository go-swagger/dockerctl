#!/bin/sh

# Execute all tests in. This is used in CI
export CI=1

find . -type f -iname "*_test.sh"|while read fname; do
    echo "Run: $fname"
    res=$(bash -c $fname 2>&1)
    if [ "$?" = "1" ]; then
        echo "$res"
        echo "Fail: $fname"
        exit 1
    fi
    echo "Pass: $fname"
done