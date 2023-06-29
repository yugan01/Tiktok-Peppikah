package main

import (
	"fmt"
	echo "orbital_demo/kitex_gen/echo"
	"context"

)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *echo.Request) (resp *echo.Response, err error) {
	fmt.Println("kitex running for service Echo")
	resp = &echo.Response{Message: req.Message}
	return
}
