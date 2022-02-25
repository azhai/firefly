package services

import (
	"context"
	"fmt"

	"github.com/astro-bug/gondor/webapi/config"
	"github.com/astro-bug/gondor/webapi/config/dialect"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
)

const RPCX_DEFAULT_PORT uint16 = 8972

var (
	rpcxProto = "tcp"
	rpcxAddr  string
	disco     client.ServiceDiscovery
)

func Initialize(cfg *config.Settings, verbose bool) {
	if len(cfg.MicroServices) > 0 {
		msrv := cfg.MicroServices[0]
		rpcxProto = msrv.Protocol
		rpcxAddr = msrv.Params.GetAddr("127.0.0.1", RPCX_DEFAULT_PORT)
	} else {
		params := new(dialect.ConnParams)
		rpcxAddr = params.GetAddr("127.0.0.1", RPCX_DEFAULT_PORT)
	}
	s := server.NewServer()
	RegisterAll(s)
	// 放在后台运行，避免阻塞主进程
	go s.Serve(rpcxProto, rpcxAddr)
}

// 注册服务，插件必须在服务对象之前注册
func RegisterAll(s *server.Server) {
	s.RegisterName("Arith", new(Arith), "")
}

func CallMethod(class, method string, args, reply interface{}) (err error) {
	if disco == nil {
		discoAddr := fmt.Sprintf("%s@%s", rpcxProto, rpcxAddr)
		disco, _ = client.NewPeer2PeerDiscovery(discoAddr, "")
	}
	c := client.NewXClient(class, client.Failtry, client.RandomSelect, disco, client.DefaultOption)
	defer c.Close()
	err = c.Call(context.Background(), method, args, reply)
	return
}

//func Test() {
//	args := &Args{
//		A: 10,
//		B: 20,
//	}
//	reply := &Reply{}
//	err := CallMethod("Arith", "Mul", args, reply)
//	if err != nil {
//		fmt.Errorf("failed to call: %v", err)
//	}
//	fmt.Printf("%d * %d = %d", args.A, args.B, reply.C)
//}
