// Copyright 2019 Andy Pan. All rights reserved.
// Copyright 2017 Joshua J Baker. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/panjf2000/gnet"
)

type echoServer struct {
	*gnet.EventServer
}

func (es *echoServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Printf("Echo server is listening on %s (multi-cores: %t, loops: %d)\n",
		srv.Addr.String(), srv.Multicore, srv.NumEventLoop)
	return
}
// func (es *echoServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
// 	fmt.Printf("Conn: %s is closed.\n", c.RemoteAddr().String())
// 	return
// }
func (es *echoServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	// Echo synchronously.
	out = frame
	return

	/*
		// Echo asynchronously.
		data := append([]byte{}, frame...)
		go func() {
			time.Sleep(time.Second)
			c.AsyncWrite(data)
		}()
		return
	*/
}

func main() {
	var port int
	var multicore bool

	// Example command: go run echo.go --port 9000 --multicore=true
	flag.IntVar(&port, "port", 9000, "--port 9000")
	flag.BoolVar(&multicore, "multicore", false, "--multicore true")
	flag.Parse()
	echo := new(echoServer)
	log.Fatal(gnet.Serve(echo, fmt.Sprintf("tcp://:%d", port), gnet.WithMulticore(multicore)))
}
