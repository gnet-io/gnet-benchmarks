package main

import (
	"log"
	"time"

	"github.com/panjf2000/ants/v2"
	"github.com/panjf2000/gnet"
)

type echoServer struct {
	*gnet.EventServer
	pool *ants.Pool
}

func (es *echoServer) React(c gnet.Conn) (out []byte, action gnet.Action) {
	data := c.ReadBytes()
	c.ResetBuffer()

	// Use ants pool to unblock the event-loop.
	_ = es.pool.Submit(func() {
		time.Sleep(1 * time.Second)
		c.AsyncWrite(data)
	})

	action = gnet.DataRead
	return
}

func main() {
	// Create a goroutine pool.
	poolSize := 64 * 1024
	pool, _ := ants.NewPool(poolSize, ants.WithNonblocking(true))
	defer pool.Release()
	echo := &echoServer{pool: pool}
	log.Fatal(gnet.Serve(echo, "tcp://:9000", gnet.WithMulticore(true)))
}
