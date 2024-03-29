// Code generated by Kitex v0.6.1. DO NOT EDIT.

package mathsvc

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	math "orbital_demo/kitex_gen/math"
)

func serviceInfo() *kitex.ServiceInfo {
	return mathSvcServiceInfo
}

var mathSvcServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "MathSvc"
	handlerType := (*math.MathSvc)(nil)
	methods := map[string]kitex.MethodInfo{
		"add":      kitex.NewMethodInfo(addHandler, newMathSvcAddArgs, newMathSvcAddResult, false),
		"subtract": kitex.NewMethodInfo(subtractHandler, newMathSvcSubtractArgs, newMathSvcSubtractResult, false),
		"multiply": kitex.NewMethodInfo(multiplyHandler, newMathSvcMultiplyArgs, newMathSvcMultiplyResult, false),
		"divide":   kitex.NewMethodInfo(divideHandler, newMathSvcDivideArgs, newMathSvcDivideResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "math",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func addHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*math.MathSvcAddArgs)
	realResult := result.(*math.MathSvcAddResult)
	success, err := handler.(math.MathSvc).Add(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMathSvcAddArgs() interface{} {
	return math.NewMathSvcAddArgs()
}

func newMathSvcAddResult() interface{} {
	return math.NewMathSvcAddResult()
}

func subtractHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*math.MathSvcSubtractArgs)
	realResult := result.(*math.MathSvcSubtractResult)
	success, err := handler.(math.MathSvc).Subtract(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMathSvcSubtractArgs() interface{} {
	return math.NewMathSvcSubtractArgs()
}

func newMathSvcSubtractResult() interface{} {
	return math.NewMathSvcSubtractResult()
}

func multiplyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*math.MathSvcMultiplyArgs)
	realResult := result.(*math.MathSvcMultiplyResult)
	success, err := handler.(math.MathSvc).Multiply(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMathSvcMultiplyArgs() interface{} {
	return math.NewMathSvcMultiplyArgs()
}

func newMathSvcMultiplyResult() interface{} {
	return math.NewMathSvcMultiplyResult()
}

func divideHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*math.MathSvcDivideArgs)
	realResult := result.(*math.MathSvcDivideResult)
	success, err := handler.(math.MathSvc).Divide(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMathSvcDivideArgs() interface{} {
	return math.NewMathSvcDivideArgs()
}

func newMathSvcDivideResult() interface{} {
	return math.NewMathSvcDivideResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Add(ctx context.Context, req *math.AddRequest) (r *math.AddResponse, err error) {
	var _args math.MathSvcAddArgs
	_args.Req = req
	var _result math.MathSvcAddResult
	if err = p.c.Call(ctx, "add", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Subtract(ctx context.Context, req *math.SubtractRequest) (r *math.SubtractResponse, err error) {
	var _args math.MathSvcSubtractArgs
	_args.Req = req
	var _result math.MathSvcSubtractResult
	if err = p.c.Call(ctx, "subtract", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Multiply(ctx context.Context, req *math.MultiplyRequest) (r *math.MultiplyResponse, err error) {
	var _args math.MathSvcMultiplyArgs
	_args.Req = req
	var _result math.MathSvcMultiplyResult
	if err = p.c.Call(ctx, "multiply", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Divide(ctx context.Context, req *math.DivideRequest) (r *math.DivideResponse, err error) {
	var _args math.MathSvcDivideArgs
	_args.Req = req
	var _result math.MathSvcDivideResult
	if err = p.c.Call(ctx, "divide", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
