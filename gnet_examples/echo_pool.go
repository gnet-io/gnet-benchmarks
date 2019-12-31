package main

import (
	"log"
	"time"

	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
)

type echoServer struct {
	*gnet.EventServer
	pool *goroutine.Pool
}

func (es *echoServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	data := append([]byte{}, frame...)

	// Use ants pool to unblock the event-loop.
	_ = es.pool.Submit(func() {
		time.Sleep(1 * time.Second)
		c.AsyncWrite(data)
	})

	return
}

func main() {
	p := goroutine.Default()
	defer p.Release()

	echo := &echoServer{pool: p}
	log.Fatal(gnet.Serve(echo, "tcp://:9000", gnet.WithMulticore(true)))
}