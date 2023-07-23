// Code generated by hertz generator. DO NOT EDIT.

package math

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	math "orbital_demo/biz/handler/math"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_math := root.Group("/math", _mathMw()...)
		_math.POST("/add", append(_addMw(), math.Add)...)
		_math.POST("/divide", append(_divideMw(), math.Divide)...)
		_math.POST("/multiply", append(_multiplyMw(), math.Multiply)...)
		_math.POST("/subtract", append(_subtractMw(), math.Subtract)...)
	}
}
