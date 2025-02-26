package main

import (
	"log"
	"net"

	pb "github.com/ttnsgn/grpc-go-handson/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Println("Listen on: %s\n", addr)
	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
