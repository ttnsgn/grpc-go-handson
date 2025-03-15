package main

import (
	"context"
	"io"
	"log"

	pb "github.com/ttnsgn/grpc-go-handson/greet/proto"
)

// func doGreet(c pb.GreetServiceClient) {
// 	log.Println("doGreet was invoked")
// 	res, err := c.Greet(context.Background(), &pb.GreetRequest{
// 		FirstName: "Tatan",
// 	})
// 	if err != nil {
// 		log.Fatalf("Could not greet: %v\n", err)
// 	}
//
// 	log.Printf("Greeting: %s\n", res.Result)
// }

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")
	stream, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{FirstName: "Tatan"})
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}
