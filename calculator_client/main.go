package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	pb "github.com/Ziayaya/grpc-posttest/calculatorpb"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

// argParser parses and converts command-line arguments to integers.
func argParser(n1, n2 string) (int32, int32) {
	N1, err := strconv.Atoi(n1)
	if err != nil {
		log.Fatalf("Cannot parse n1: %s", err)
	}
	N2, err := strconv.Atoi(n2)
	if err != nil {
		log.Fatalf("Cannot parse n2: %s", err)
	}
	return int32(N1), int32(N2)
}

func main() {
	// Ensure that exactly 2 command-line arguments are provided.
	if len(os.Args) != 3 {
		log.Fatalf("2 numbers expected: n1 n2")
	}

	// Parse command-line arguments.
	n1, n2 := argParser(os.Args[1], os.Args[2])

	// Establish a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect: %s", err)
	}
	defer conn.Close()

	// Create a gRPC client.
	client := pb.NewCalculatorClient(conn)

	// Set a timeout for gRPC requests.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Perform gRPC Add operation.
	addResult, err := client.Add(ctx, &pb.AddRequest{N1: int32(n1), N2: int32(n2)})
	if err != nil {
		log.Fatalf("Adding error: %s", err)
	}
	log.Printf("%d + %d = %d", n1, n2, addResult.N1)

	// Perform gRPC Subtract operation.
	subtractResult, err := client.Subtract(ctx, &pb.SubtractRequest{N1: int32(n1), N2: int32(n2)})
	if err != nil {
		log.Fatalf("Subtracting error: %s", err)
	}
	log.Printf("%d - %d = %d", n1, n2, subtractResult.N1)

	// Perform gRPC Multiply operation.
	multiplyResult, err := client.Multiply(ctx, &pb.MultiplyRequest{N1: int32(n1), N2: int32(n2)})
	if err != nil {
		log.Fatalf("Multiplying error: %s", err)
	}
	log.Printf("%d * %d = %d", n1, n2, multiplyResult.N1)

	// Perform gRPC Divide operation.
	divideResult, err := client.Divide(ctx, &pb.DivideRequest{N1: int32(n1), N2: int32(n2)})
	if err != nil {
		// Check if the error is due to division by zero.
		if strings.Contains(err.Error(), "Cannot divide by zero") {
			log.Printf("Dividing error: Cannot divide by zero")
		} else {
			log.Fatalf("Dividing error: %s", err)
		}
	} else {
		// Display the result with two decimal places.
		result := float64(divideResult.N1)
		log.Printf("%d / %d = %.2f", n1, n2, result)
	}
}
