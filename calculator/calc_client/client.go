package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tokada11/grpc-go-course/calculator/calcpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calcpb.NewCalcServiceClient(cc)
	//fmt.Printf("Created client %f", c)
	doUnary(c)
}

func doUnary(c calcpb.CalcServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")

	req := &calcpb.CalcRequest{
		Numbers: &calcpb.Numbers{
			FirstNum:  3,
			SecondNum: 10,
		},
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Sum: %v", res.Answer)
}
