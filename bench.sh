#!/bin/bash

set -e

cd $(dirname "${BASH_SOURCE[0]}")

if [ ! -d "result/" ];then
    mkdir -p result/
fi

./bench-http.sh 2>&1 | tee result/http.txt
./bench-echo.sh 2>&1 | tee result/echo.txt

go run analyze.go