package main

import (
	"context"
	"log"

	pb "github.com/ttnsgn/grpc-go-handson/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", in)
	var result int32 = in.FisrtNumber + in.SecondNumer
	return &pb.SumResponse{
		Result: result,
	}, nil
}

func (s *Server) Prime(in *pb.PrimeRequest, stream grpc.ServerStreamingServer[pb.PrimeResponse]) error {
	log.Printf("Prime function was invoked with %v\n", in)
	var k int32 = 2
	var n int32 = in.InputNumber
	for n > 1 {
		if n%k == 0 {
			err := stream.Send(&pb.PrimeResponse{Result: k})
			if err != nil {
				log.Fatalf("Error while send stream %v\n", err)
			}
			n = n / k
		} else {
			k = k + 1
		}
	}
	return nil
}
