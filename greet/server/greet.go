package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ttnsgn/grpc-go-handson/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream grpc.ServerStreamingServer[pb.GreetResponse]) error {
	log.Printf("GreetManyTimes was invoked with: %v\n", in)
	for i := range 10 {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)
		err := stream.Send(&pb.GreetResponse{
			Result: res,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
