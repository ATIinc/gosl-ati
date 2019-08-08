#!/bin/bash

GP1="${GOPATH%:*}"
GP2="${GOPATH#*:}"
GP=$GP2
if [[ -z "${GP// }" ]]; then
    GP=$GP1
fi
GOSL="$GP/pkg/linux_amd64/github.com/cpmech/gosl/gm.a"

FILES="*.go"

if [ -f $GOSL ]; then
   FILES="$FILES $GOSL"
fi

echo
echo "monitoring:"
echo $FILES
echo
echo "with:"
echo "GP = $GP"
echo
echo

while true; do
    inotifywait -q -e modify $FILES
    echo
    echo
    echo
    echo
    #go run h5_ang_mnist01.go
    go run ml_simple01.go
    go run ml_ang01.go
    go run ml_ang02.go
done
