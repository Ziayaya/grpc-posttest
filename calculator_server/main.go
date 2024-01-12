package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Ziayaya/grpc-posttest/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
)

const (
	protocol = "tcp"
	port     = ":50051"
)

// server is used to implement the Calculator service.
type server struct {
	pb.UnimplementedCalculatorServer
}

// Add implements the Add RPC method of the Calculator service.
func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddReply, error) {
	log.Printf("Received request to add with N1 = %d and N2 = %d\n", in.N1, in.N2)
	return &pb.AddReply{N1: in.N1 + in.N2}, nil
}

// Subtract implements the Subtract RPC method of the Calculator service.
func (s *server) Subtract(ctx context.Context, in *pb.SubtractRequest) (*pb.SubtractReply, error) {
	log.Printf("Received request to subtract with N1 = %d and N2 = %d\n", in.N1, in.N2)
	return &pb.SubtractReply{N1: in.N1 - in.N2}, nil
}

// Multiply implements the Multiply RPC method of the Calculator service.
func (s *server) Multiply(ctx context.Context, in *pb.MultiplyRequest) (*pb.MultiplyReply, error) {
	log.Printf("Received request to multiply with N1 = %d and N2 = %d\n", in.N1, in.N2)
	return &pb.MultiplyReply{N1: in.N1 * in.N2}, nil
}

// Divide implements the Divide RPC method of the Calculator service.
func (s *server) Divide(ctx context.Context, in *pb.DivideRequest) (*pb.DivideReply, error) {
	log.Printf("Received request to divide with N1 = %d and N2 = %d\n", in.N1, in.N2)
	if in.N2 == 0 {
		return nil, grpc.Errorf(codes.InvalidArgument, "Cannot divide by zero")
	}
	return &pb.DivideReply{N1: in.N1 / in.N2}, nil
}

func main() {
	// Listen for incoming connections.
	lis, err := net.Listen(protocol, port)
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	// Create a new gRPC server.
	s := grpc.NewServer()

	// Register the Calculator service with the server.
	pb.RegisterCalculatorServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	// Serve and handle incoming requests.
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
