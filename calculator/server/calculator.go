package main

import (
	"context"
	"log"

	pb "github.com/ttnsgn/grpc-go-handson/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Println("Sum function was invoked with %v\n", in)
	var result int32 = in.FisrtNumber + in.SecondNumer
	return &pb.SumResponse{
		Result: result,
	}, nil
}
