package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/tokada11/grpc-go-course/calculator/calcpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calcpb.CalcRequest) (*calcpb.CalcResponse, error) {
	fmt.Printf("Sum function was invoked with %v", req)
	firstNum := req.Numbers.FirstNum
	secondNum := req.Numbers.SecondNum
	answer := firstNum + secondNum
	res := &calcpb.CalcResponse{
		Answer: answer,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello Calculator Service")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalcServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
