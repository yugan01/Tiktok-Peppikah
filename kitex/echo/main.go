package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/registry-nacos/registry"
	echo "orbital_demo/kitex_gen/echo/echo"
	"github.com/cloudwego/kitex/server"
)

func main() {
	r, err := registry.NewDefaultNacosRegistry()
	if err != nil {
		log.Println("Nacos Registry error:", err)
	}

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8001")
	if err != nil {
		panic(err)
	}

	svr := echo.NewServer(new(EchoImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "echo"}),
		server.WithRegistry(r),
		server.WithServiceAddr(addr),
	)

	if err := svr.Run(); err != nil {
		log.Println("Server stopped with error:", err)
	} else {
		log.Println("Server has stopped")
	}
}
