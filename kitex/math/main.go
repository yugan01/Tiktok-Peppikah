package main

import (
	"log"
	"net"
	math "orbital_demo/kitex_gen/math/mathsvc"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
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

	svr := math.NewServer(new(MathSvcImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "math"}),
		server.WithRegistry(r),
		server.WithServiceAddr(addr),
	)

	if err := svr.Run(); err != nil {
		log.Println("Server stopped with error:", err)
	} else {
		log.Println("Server has stopped")
	}
}
