package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/panjf2000/gnet/v2"
)

var (
	errMsg      = "Internal Server Error"
	errMsgBytes = []byte(errMsg)
)

type httpServer struct {
	*gnet.BuiltinEventEngine

	parser    httpParser
	addr      string
	multicore bool
	eng       gnet.Engine
}

type httpParser struct {
	delimiter []byte

	rsp []byte
}

func (p *httpParser) Parse(c gnet.Conn) error {
	buf, _ := c.Next(-1)

	// process the pipeline
	var i int
pipeline:
	if i = bytes.Index(buf, p.delimiter); i != -1 {
		p.rsp = append(p.rsp, "HTTP/1.1 200 OK\r\nServer: gnet\r\nContent-Type: text/plain\r\nDate: "...)
		p.rsp = time.Now().AppendFormat(p.rsp, "Mon, 02 Jan 2006 15:04:05 GMT")
		p.rsp = append(p.rsp, "\r\nContent-Length: 13\r\n\r\nHello, World!"...)
		buf = buf[i+4:]
		goto pipeline
	}

	// request not ready, yet
	return nil
}

func (hs *httpServer) OnBoot(eng gnet.Engine) gnet.Action {
	hs.eng = eng
	log.Printf("echo server with multi-core=%t is listening on %s\n", hs.multicore, hs.addr)
	return gnet.None
}

func (hs *httpServer) OnTraffic(c gnet.Conn) gnet.Action {
	if hs.parser.Parse(c) != nil {
		// bad thing happened
		c.Write(errMsgBytes)
		return gnet.Close
	}

	// handle the request
	if len(hs.parser.rsp) > 0 {
		c.Write(hs.parser.rsp)
		hs.parser.rsp = hs.parser.rsp[:0]
	}
	return gnet.None
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
}

func main() {
	var port int
	var multicore bool

	// Example command: go run main.go --port 8080 --multicore=true
	flag.IntVar(&port, "port", 8080, "server port")
	flag.BoolVar(&multicore, "multicore", true, "multicore")
	flag.Parse()

	hc := httpParser{delimiter: []byte("\r\n\r\n")}
	http := &httpServer{addr: fmt.Sprintf("tcp://127.0.0.1:%d", port), multicore: multicore, parser: hc}

	// Start serving!
	log.Println("server exits:", gnet.Run(http, http.addr, gnet.WithMulticore(multicore)))
}
