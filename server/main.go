// Server application that serves mathematical calculations via gRPC
package main

import (
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/rolandktud/calculator/calculatorService"
	server "github.com/rolandktud/calculator/server/serverImpl"
)

var (
	address = flag.String("address", "localhost", "address of the calculator server")
	port    = flag.String("port", "30000", "port of the calculator server")
)

func main() {
	lis, err := net.Listen("tcp", (*address)+":"+(*port))
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server.Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
