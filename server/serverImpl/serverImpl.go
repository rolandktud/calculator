// Implementation for gRPC server methods
package serverImpl

import (
	"context"
	"errors"

	pb "github.com/rolandktud/calculator/calculatorService"
)

type Server struct{}

// serve addition request
func (s *Server) Addition(ctx context.Context, in *pb.Operands) (*pb.CalculationReply, error) {
	result, err := internalAddition(in.A, in.B)
	return &pb.CalculationReply{Result: result}, err
}

// serve substraction request
func (s *Server) Subtraction(ctx context.Context, in *pb.Operands) (*pb.CalculationReply, error) {
	result, err := internalSubstraction(in.A, in.B)
	return &pb.CalculationReply{Result: result}, err
}

// serve multiplication request
func (s *Server) Multiplication(ctx context.Context, in *pb.Operands) (*pb.CalculationReply, error) {
	result, err := internalMultiplication(in.A, in.B)
	return &pb.CalculationReply{Result: result}, err
}

// serve division request
func (s *Server) Division(ctx context.Context, in *pb.Operands) (*pb.CalculationReply, error) {
	result, err := internalDivision(in.A, in.B)
	return &pb.CalculationReply{Result: result}, err
}

// internal addition
func internalAddition(a float64, b float64) (float64, error) {
	return a + b, nil
}

// internal substraction
func internalSubstraction(a float64, b float64) (float64, error) {
	return a - b, nil
}

// internal multiplication
func internalMultiplication(a float64, b float64) (float64, error) {
	return a * b, nil
}

// internal division
func internalDivision(a float64, b float64) (float64, error) {
	if b != 0 {
		return a / b, nil
	}
	return 0, errors.New("illegal division by 0")

}
