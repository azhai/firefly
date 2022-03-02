package main

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/cloudwego/netpoll"
)

const RelayKey = "relay"
var relayPool = sync.Pool {
    New: func() any {
		relay, _ := netpoll.DialConnection("tcp", "127.0.0.1:6379", time.Second)
		// relay.SetOnRequest(handle)
		return relay
    },
}

func main() {
	listener, err := netpoll.CreateListener("tcp", ":6380")
	if err != nil {
		panic("create netpoll listener failed")
	}

	eventLoop, _ := netpoll.NewEventLoop(
		handle,
		netpoll.WithOnPrepare(prepare),
		netpoll.WithReadTimeout(time.Second),
	)
	// start listen loop ...
	eventLoop.Serve(listener)

	// stop server ...
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	eventLoop.Shutdown(ctx)
}

func prepare(conn netpoll.Connection) (ctx context.Context) {
	conn.AddCloseCallback(finish)
	relay := relayPool.Get().(netpoll.Connection)
	ctx = context.WithValue(context.Background(), RelayKey, relay)
	return
}

func finish(conn netpoll.Connection) (err error) {
	ctx := context.Background()
	relay := ctx.Value(RelayKey).(netpoll.Connection)
	relayPool.Put(relay)
	return
}

func handle(ctx context.Context, conn netpoll.Connection) (err error) {
	relay := ctx.Value(RelayKey).(netpoll.Connection)
	// go io.Copy(conn, relay) // 复制服务端回应
	// NOTICE: 与上面一行不能对调，否则无法知道客户端关闭了
	// _, err = io.Copy(relay, conn) // 复制上报数据
	return nil
}
