#!/bin/bash

FILE="*.go *.c *.f"

while true; do
    inotifywait -q -e modify $FILE
    echo
    echo
    echo
    echo
    go test -test.run="Chebmo"
done
