#!/bin/bash

set -e

echo ""
echo "--- BENCH ECHO START ---"
echo ""

echo "*** 100 connections, 10 seconds, 6 byte packets"
nl=$'\r\n'
tcpkali --workers 4 -c 100 -T 10s -m "PING{$nl}" 127.0.0.1:5000
echo "--- DONE ---"

