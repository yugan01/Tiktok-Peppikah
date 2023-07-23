package main

import (
	"context"
	"fmt"
	math "orbital_demo/kitex_gen/math"
)

// MathSvcImpl implements the last service interface defined in the IDL.
type MathSvcImpl struct{}

// Add implements the MathSvcImpl interface.
func (s *MathSvcImpl) Add(ctx context.Context, req *math.AddRequest) (resp *math.AddResponse, err error) {
	fmt.Println("kitex running for service Math/Add")
	resp = &math.AddResponse{Addresult: req.First + req.Second}
	return
}

// Subtract implements the MathSvcImpl interface.
func (s *MathSvcImpl) Subtract(ctx context.Context, req *math.SubtractRequest) (resp *math.SubtractResponse, err error) {
	fmt.Println("kitex running for service Math/Subtract")
	resp = &math.SubtractResponse{Subtractresult: req.First - req.Second}
	return
}

// Multiply implements the MathSvcImpl interface.
func (s *MathSvcImpl) Multiply(ctx context.Context, req *math.MultiplyRequest) (resp *math.MultiplyResponse, err error) {
	fmt.Println("kitex running for service Math/Multiply")
	resp = &math.MultiplyResponse{Multiplyresult: req.First * req.Second}
	return
}

// Divide implements the MathSvcImpl interface.
func (s *MathSvcImpl) Divide(ctx context.Context, req *math.DivideRequest) (resp *math.DivideResponse, err error) {
	fmt.Println("kitex running for service Math/Divide")
	resp = &math.DivideResponse{Divideresult: req.First / req.Second}
	return
}
