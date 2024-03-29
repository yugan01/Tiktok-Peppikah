// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/pprof"
)

func main() {
	
	//Starting the server
	h := server.Default()

	//Using Pprof for graphical analysis of this project
	pprof.Register(h, "dev/pprof")

	register(h)
	h.Spin()
}
