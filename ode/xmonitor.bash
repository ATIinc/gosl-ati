#!/bin/bash

FILE="*.go"

while true; do
    inotifywait -q -e modify $FILE
    echo
    echo
    echo
    echo
    #go test -test.run="Ode01"
    #go test -test.run="BwEuler"
    #go test -test.run="Radau502"
    #go test -test.run="DoPri502"
    #go test -test.run="Erk05"
    go test -test.run="HL01"
done
