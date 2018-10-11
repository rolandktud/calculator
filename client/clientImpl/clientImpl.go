package clientImpl

import (
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc"

	pb "github.com/rolandktud/calculator/calculatorService"
)

type Client struct {
	clientImpl pb.CalculatorClient
}

// NewClient creates a new client using an existing connection
func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{pb.NewCalculatorClient(conn)}
}

// Calculate connects to the gRPC server for performing the calculations
func (c *Client) Calculate(operation string, operand_a float64, operand_b float64) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var (
		r   *pb.CalculationReply
		err error
	)
	switch operation {
	case "add":
		r, err = c.clientImpl.Addition(ctx, &pb.Operands{A: operand_a, B: operand_b})
	case "sub":
		r, err = c.clientImpl.Subtraction(ctx, &pb.Operands{A: operand_a, B: operand_b})
	case "mul":
		r, err = c.clientImpl.Multiplication(ctx, &pb.Operands{A: operand_a, B: operand_b})
	case "div":
		r, err = c.clientImpl.Division(ctx, &pb.Operands{A: operand_a, B: operand_b})
	default:
		err = errors.New(fmt.Sprintf("Operation not found: %s", operation))
	}

	if err != nil {
		return 0, err
	}

	if r == nil {
		panic("result was 0 but no error code returned")
	}

	return r.Result, nil
}
