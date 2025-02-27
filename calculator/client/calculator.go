package main

import (
	"context"
	"io"
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

func doPrime(c pb.CalculatorServiceClient) {
	log.Println("func doPrime was invoked")
	stream, err := c.Prime(context.Background(), &pb.PrimeRequest{InputNumber: 120})
	if err != nil {
		log.Fatalf("Error while calling Prime stream %v\n", err)
	}

	primes := []int32{}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream %v\n", err)
		}

		primes = append(primes, msg.Result)
	}

	log.Printf("Primes number from 120 is %v", primes)
}
