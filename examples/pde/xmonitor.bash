#!/bin/bash

GP1="${GOPATH%:*}"
GP2="${GOPATH#*:}"
GP=$GP2
if [[ -z "${GP// }" ]]; then
    GP=$GP1
fi
GOSL="$GP/pkg/linux_amd64/github.com/cpmech/gosl/pde.a"

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
    go run spc_fdm_simple01.go
done
