#!/bin/bash

set -e

echo ""
echo "--- BENCH ECHO START ---"
echo ""

cd $(dirname "${BASH_SOURCE[0]}")
function cleanup() {
  echo "--- BENCH ECHO DONE ---"
  kill -9 $(jobs -rp)
  wait $(jobs -rp) 2>/dev/null
}
trap cleanup EXIT

mkdir -p bin
$(pkill -9 net-echo-server || printf "")
$(pkill -9 evio-echo-server || printf "")
$(pkill -9 eviop-echo-server || printf "")
$(pkill -9 gev-echo-server || printf "")
$(pkill -9 gnet-echo-server || printf "")

conn_num=$1
test_duration=$2
packet_size=$3
data=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w $packet_size | head -n 1)
echo "--- ECHO PACKET ---"
echo $data
echo ""

function gobench() {
  echo "--- $1 ---"
  echo ""
  if [[ "$1" == "GNET" ]]; then
    go build -tags=poll_opt -gcflags="-l=4" -ldflags="-s -w" -o $2 $3
  else
    go build -gcflags="-l=4" -ldflags="-s -w" -o $2 $3
  fi

  if [[ "$1" == "GO-NET" ]]; then
    $2 --port $4 &
  elif [[ "$1" == "GNET" ]]; then
    $2 --port $4 --multicore=$5 &
  else
    $2 --port $4 --loops $5 &
  fi

  sleep 1
  printf "*** %d connections, %d seconds, packet size: %d bytes\n" $conn_num $test_duration $packet_size
  
  nl=$'\r\n'
  tcpkali -c $conn_num -T $test_duration's' -m $data 127.0.0.1:$4
  echo ""
  echo "--- DONE ---"
  echo ""
}

gobench "GO-NET" bin/net-echo-server net-echo-server/main.go 5000
gobench "EVIO" bin/evio-echo-server evio-echo-server/main.go 5001 -1
gobench "GEV" bin/gev-echo-server gev-echo-server/echo.go 5003 -1
gobench "GNET" bin/gnet-echo-server gnet-echo-server/main.go 5004 true
