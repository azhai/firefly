package main

import (
	"context"
	"time"

	"github.com/cloudwego/netpoll"
)

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
	return
}

func connect(ctx context.Context, conn netpoll.Connection) context.Context {
	return ctx
}

func finish(conn netpoll.Connection) (err error) {
	return
}

func handle(ctx context.Context, conn netpoll.Connection) (err error) {
	return nil
}
