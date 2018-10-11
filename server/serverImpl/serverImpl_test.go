// Tests for implementation of the gRPC client for calculator service
package serverImpl

import (
	"testing"

	pb "github.com/rolandktud/calculator/calculatorService"
)

// testing correct additions
func TestAddition(t *testing.T) {
	var tests = []struct {
		input pb.Operands
		want  pb.CalculationReply
	}{
		{pb.Operands{A: 1.0, B: 2.0}, pb.CalculationReply{Result: 3.0}},
		{pb.Operands{A: 3.0, B: -5.0}, pb.CalculationReply{Result: -2.0}},
	}

	server := Server{}

	for _, test := range tests {
		if got, err := server.Addition(nil, &test.input); *got != test.want || err != nil {
			t.Errorf(" server.Addition(nil, &Operands%v) = %v", test.input, got)
		}
	}
}

// testing correct substractions
func TestSubstraction(t *testing.T) {
	var tests = []struct {
		input pb.Operands
		want  pb.CalculationReply
	}{
		{pb.Operands{A: 1.0, B: 2.0}, pb.CalculationReply{Result: -1.0}},
		{pb.Operands{A: 3.0, B: -5.0}, pb.CalculationReply{Result: 8.0}},
	}

	server := Server{}

	for _, test := range tests {
		if got, err := server.Subtraction(nil, &test.input); *got != test.want || err != nil {
			t.Errorf(" server.Subtraction(nil, &Operands%v) = %v", test.input, got)
		}
	}
}

// testing correct multiplications
func TestMultiplication(t *testing.T) {
	var tests = []struct {
		input pb.Operands
		want  pb.CalculationReply
	}{
		{pb.Operands{A: 1.0, B: 2.0}, pb.CalculationReply{Result: 2.0}},
		{pb.Operands{A: 3.0, B: -5.0}, pb.CalculationReply{Result: -15.0}},
	}

	server := Server{}

	for _, test := range tests {
		if got, err := server.Multiplication(nil, &test.input); *got != test.want || err != nil {
			t.Errorf(" server.Multiplication(nil, &Operands%v) = %v", test.input, got)
		}
	}
}

// testing correct divisions
func TestDivision(t *testing.T) {
	var tests = []struct {
		input pb.Operands
		want  pb.CalculationReply
	}{
		{pb.Operands{A: 1.0, B: 2.0}, pb.CalculationReply{Result: 0.5}},
		{pb.Operands{A: 3.0, B: -5.0}, pb.CalculationReply{Result: -0.6}},
	}

	server := Server{}

	for _, test := range tests {
		if got, err := server.Division(nil, &test.input); *got != test.want || err != nil {
			t.Errorf(" server.Division(nil, &Operands%v) = %v", test.input, got)
		}
	}

	// testing division by 0
	divisionBy0Test := pb.Operands{A: 1.0, B: 0.0}
	got, err := server.Division(nil, &divisionBy0Test)
	if err == nil {
		t.Errorf("expected error in division with 0 server.Division(nil, &Operands%v) = %v", divisionBy0Test, got)
	}
}
