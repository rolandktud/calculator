// Tests for implementation of the gRPC client for calculator service
package clientImpl

import (
	pb "calculator/calculatorService"
	"testing"

	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Mock calculator client for checking if certain callbacks are triggered
type mockCalculatorClient struct {
	calledAddition       bool
	calledSubtraction    bool
	calledMultiplication bool
	calledDivision       bool
}

func (m *mockCalculatorClient) Addition(ctx context.Context, in *pb.Operands, opts ...grpc.CallOption) (*pb.CalculationReply, error) {
	m.calledAddition = true
	return &pb.CalculationReply{Result: 0}, nil
}
func (m *mockCalculatorClient) Subtraction(ctx context.Context, in *pb.Operands, opts ...grpc.CallOption) (*pb.CalculationReply, error) {
	m.calledSubtraction = true
	return &pb.CalculationReply{Result: 0}, nil
}
func (m *mockCalculatorClient) Multiplication(ctx context.Context, in *pb.Operands, opts ...grpc.CallOption) (*pb.CalculationReply, error) {
	m.calledMultiplication = true
	return &pb.CalculationReply{Result: 0}, nil
}
func (m *mockCalculatorClient) Division(ctx context.Context, in *pb.Operands, opts ...grpc.CallOption) (*pb.CalculationReply, error) {
	m.calledDivision = true
	return &pb.CalculationReply{Result: 0}, nil
}

func TestCalculate(t *testing.T) {

	// test case: correct addition
	mock := mockCalculatorClient{}
	client := Client{&mock}
	client.Calculate("add", 1, 2)
	if !mock.calledAddition {
		t.Errorf("client.Calculate(%s, %d, %d) did not call Addition procedure", "add", 1, 2)
	}

	// test case: correct substraction
	mock = mockCalculatorClient{}
	client = Client{&mock}
	client.Calculate("sub", 1, 2)
	if !mock.calledSubtraction {
		t.Errorf("client.Calculate(%s, %d, %d) did not call Substraction procedure", "add", 1, 2)
	}

	// test case: correct multiplication
	mock = mockCalculatorClient{}
	client = Client{&mock}
	client.Calculate("mul", 1, 2)
	if !mock.calledMultiplication {
		t.Errorf("client.Calculate(%s, %d, %d) did not call Multiplication procedure", "add", 1, 2)
	}

	// test case: correct division
	mock = mockCalculatorClient{}
	client = Client{&mock}
	client.Calculate("div", 1, 2)
	if !mock.calledDivision {
		t.Errorf("client.Calculate(%s, %d, %d) did not call Division procedure", "add", 1, 2)
	}

	// test case: don't call any routine with invalid code
	mock = mockCalculatorClient{}
	client = Client{&mock}
	client.Calculate("__", 1, 2)
	if mock.calledAddition || mock.calledSubtraction || mock.calledMultiplication || mock.calledDivision {
		t.Errorf("client.Calculate(%s, %d, %d) called procedure, when invalid input was given", "__", 1, 2)
	}
}
