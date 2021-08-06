// Copyright 2017 Joshua J Baker. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 5000, "server port")
	flag.Parse()
	ln, err := net.Listen("tcp4", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	log.Printf("echo server started on port %d", port)
	var id int
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		id++
		go func(id int, conn net.Conn) {
			defer func() {
				//log.Printf("closed: %d", id)
				_ = conn.Close()
			}()
			//log.Printf("opened: %d: %s", id, conn.RemoteAddr().String())
			inBuf := make([]byte, 64*1024)
			outBuf := bytes.NewBuffer(make([]byte, 0, 64*1024))
			for {
				rn, err := conn.Read(inBuf)
				if err != nil {
					return
				}

				var wn int
				if outBuf.Len() != 0 {
					outBuf.Write(inBuf[:rn])
					wn, err = conn.Write(outBuf.Bytes())
					outBuf.Next(wn)
				} else {
					wn, err = conn.Write(inBuf[:rn])
					if wn != rn {
						outBuf.Write(inBuf[wn:])
					}
				}
				if err != nil {
					return
				}
			}
		}(id, conn)
	}
}
