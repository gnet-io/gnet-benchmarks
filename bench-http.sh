#!/bin/bash

set -e

echo ""
echo "--- BENCH HTTP START ---"
echo ""

cd $(dirname "${BASH_SOURCE[0]}")
function cleanup {
    echo "--- BENCH HTTP DONE ---"
    kill -9 $(jobs -rp)
    wait $(jobs -rp) 2>/dev/null
}
trap cleanup EXIT

mkdir -p bin
$(pkill -9 net-http-server || printf "")
$(pkill -9 fasthttp-server || printf "")
$(pkill -9 evio-http-server || printf "")
$(pkill -9 gnet-http-server || printf "")

function gobench {
    echo "--- $1 ---"
    if [[ "$3" != "" ]]; then
        go build -o $2 $3
    fi

    if [[ "$1" == "GNET" ]]; then
        $2 --port $4 --multicore=$5 &
    elif [[ "$1" == "EVIO" ]]; then
        $2 --port $4 --loops $5 &
    else
        $2 --port $4 &
    fi

    sleep 1
    echo "*** 100 connections, 10 seconds"
    # bombardier -c 100 -d 10s -l http://127.0.0.1:$4
    # wrk -c256 -d10s --latency http://127.0.0.1:$4
    wrk -H 'Host: 127.0.0.1' -H 'Accept: text/plain,text/html;q=0.9,application/xhtml+xml;q=0.9,application/xml;q=0.8,*/*;q=0.7' -H 'Connection: keep-alive' --latency -d 15 -c 256 --timeout 8 -t 4 http://127.0.0.1:$4/plaintext -s pipeline.lua -- 16
    echo "--- DONE ---"
    echo ""
}

gobench "GO-HTTP" bin/net-http-server net-http-server/main.go 8081
gobench "FASTHTTP" bin/fasthttp-server fasthttp-server/main.go 8083
gobench "EVIO" bin/evio-http-server evio-http-server/main.go 8084 -1
gobench "GNET" bin/gnet-http-server gnet-http-server/main.go 8085 true
