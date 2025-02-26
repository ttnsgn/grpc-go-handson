package main

import (
	"context"
	"log"

	pb "github.com/ttnsgn/grpc-go-handson/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("func doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FisrtNumber: 2,
		SecondNumer: 3,
	})
	if err != nil {
		log.Fatalf("Could not do sum: %v\n", err)
	}

	log.Printf("Sum result: %d", res.Result)
}
