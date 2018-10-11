// client application that takes request via STDIN and relays calculation to a remote server
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"calculator/client/clientImpl"

	"google.golang.org/grpc"
)

var (
	address = flag.String("address", "localhost", "address of the calculator server")
	port    = flag.String("port", "30000", "port of the calculator server")
)

func main() {
	flag.Parse()

	if len(flag.Args()) > 0 {
		log.Fatalf("Unkown arguments(s): %v", flag.Args())
		flag.Usage()
	}

	conn, err := grpc.Dial((*address)+":"+(*port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}
	defer conn.Close()
	c := clientImpl.NewClient(conn)

	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error while parsing input: %v. Ignoring...", err)
		}

		inputs := strings.Split(strings.Trim(input, "\n"), " ")

		if len(inputs) != 3 {
			log.Printf("incorrect parameter count:%v, exactly 3 required in form <op> <nr> <nr>, where <op> is one of [add, sub, mul, div] and <nr> is a decimal number", inputs)
			continue
		}

		operation := inputs[0]

		operand_a, err := strconv.ParseFloat(inputs[1], 64)
		if err != nil {
			log.Printf("could not parse as number %s: %v", inputs[1], err)
			continue
		}

		operand_b, err := strconv.ParseFloat(inputs[2], 64)
		if err != nil {
			log.Printf("could not parse as number %s: %v", inputs[2], err)
			continue
		}

		results, err := c.Calculate(operation, operand_a, operand_b)
		if err != nil {
			log.Printf("error during calculation: %v", err)
			continue
		}

		if isInt(results) {
			fmt.Printf("Result: %d\n", int(results))
		} else {
			fmt.Printf("Result: %f\n", results)
		}
	}
}

func isInt(val float64) bool {
	return val == float64(int(val))
}
