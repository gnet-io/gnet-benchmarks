// Copyright 2017 Joshua J Baker. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 7000, "server port")
	flag.Parse()
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	log.Printf("echo server started on port %d", port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(conn net.Conn) {
			defer conn.Close()
			var (
				rn, wn int
				buf    [0x4000]byte
			)
			for {
				rn, err = conn.Read(buf[:])
				if err != nil {
					return
				}
				wn, err = conn.Write(buf[:rn])
				if err != nil || wn != rn {
					return
				}
			}
		}(conn)
	}
}
