#!/bin/bash

set -e

echo ""
echo "--- BENCH ECHO START ---"
echo ""

cd $(dirname "${BASH_SOURCE[0]}")
function cleanup {
    echo "--- BENCH ECHO DONE ---"
    kill -9 $(jobs -rp)
    wait $(jobs -rp) 2>/dev/null
}
trap cleanup EXIT

mkdir -p bin
$(pkill -9 net-echo-server || printf "")
$(pkill -9 evio-echo-server || printf "")
$(pkill -9 gnet-echo-server || printf "")

function gobench {
    echo "--- $1 ---"
    if [[ "$3" != "" ]]; then
        go build -o $2 $3
    fi

    if [[ "$1" == "GO-NET" ]]; then
        GOMAXPROCS=4 $2 --port $4 &
    elif [[ "$1" == "GNET" ]]; then
        GOMAXPROCS=4 $2 --port $4 --multicore $5 &
    else
        GOMAXPROCS=4 $2 --port $4 --loops $5 &
    fi

    sleep 1
    echo "*** 100 connections, 10 seconds, 6 byte packets"
    nl=$'\r\n'
    tcpkali --workers 4 -c 100 -T 10s -m "PING{$nl}" 127.0.0.1:$4
    echo "--- DONE ---"
    echo ""
}

gobench "GO-NET" bin/net-echo-server net-echo-server/main.go 5000
gobench "EVIO" bin/evio-echo-server evio-echo-server/main.go 5001 -1
gobench "EVIOP" bin/eviop-echo-server eviop-echo-server/main.go 5002 -1
gobench "GEV"  bin/gev-echo-server gev-echo-server/echo.go 5003 -1
gobench "GNET" bin/gnet-echo-server gnet-echo-server/main.go 5004 true
