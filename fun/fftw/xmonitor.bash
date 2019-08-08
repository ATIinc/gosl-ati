#!/bin/bash

FILE="*.go"

while true; do
    inotifywait -q -e modify $FILE
    echo
    echo
    echo
    echo
    go test -test.run="OneDver05"
    #go test -test.run="TwoDver02"
done
