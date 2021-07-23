package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/cloudwego/netpoll"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "server port")

	network, address := "tcp", fmt.Sprintf("127.0.0.1:%d", port)

	// 创建 listener
	listener, err := netpoll.CreateListener(network, address)
	if err != nil {
		panic(fmt.Sprintf("create netpoll listener fail, error: %v", err))
	}

	// handle: 连接读数据和处理逻辑
	var onRequest netpoll.OnRequest = handler

	// options: EventLoop 初始化自定义配置项
	var opts = []netpoll.Option{
		netpoll.WithReadTimeout(1 * time.Second),
		netpoll.WithIdleTimeout(10 * time.Minute),
		netpoll.WithOnPrepare(nil),
	}

	// 创建 EventLoop
	eventLoop, err := netpoll.NewEventLoop(onRequest, opts...)
	if err != nil {
		panic(fmt.Sprintf("create netpoll event-loop fail, error: %v", err))
	}

	// 运行 Server
	err = eventLoop.Serve(listener)
	if err != nil {
		panic(fmt.Sprintf("netpoll server exit, error: %v", err))
	}
}

// 读事件处理
func handler(ctx context.Context, connection netpoll.Connection) error {
	return connection.Writer().Flush()
}
