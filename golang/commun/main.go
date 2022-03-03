package main

import (
	"context"
	"io"
	"net"
	"sync"
	"time"

	"github.com/cloudwego/netpoll"
)

const (
	RelayKey   = "relay"
	MaxPkgSize = 4096
)

var relayPool = sync.Pool{
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

	go pollCopy(conn.Writer(), relay.Reader()) // 复制服务端回应
	// NOTICE: 与上面一行不能对调，否则无法知道客户端关闭了
	_, err = pollCopy(relay.Writer(), conn.Reader()) // 复制上报数据

	return nil
}

func ioCopy(writer, reader net.Conn) (int64, error) {
	return io.Copy(writer, reader)
}

func pollCopy(writer netpoll.Writer, reader netpoll.Reader) (int64, error) {
	pipe := make(chan *netpoll.LinkBuffer)

	// writing
	go func(pipe <-chan *netpoll.LinkBuffer) {
		for {
			select {
			case pkg := <-pipe:
				writer.Append(pkg)
			default:
				if writer.MallocLen() > 0 {
					writer.Flush()
				}
			}
		}
	}(pipe)

	// reading
	for {
		pkg, _ := reader.Slice(MaxPkgSize)
		go func(pipe chan<- *netpoll.LinkBuffer, pkg netpoll.Reader) {
			buf := netpoll.NewLinkBuffer()
			data, _ := buf.ReadBinary(buf.Len())
			buf.WriteBinary(data)
			pipe <- buf
			pkg.Release()
		}(pipe, pkg)
	}
}
