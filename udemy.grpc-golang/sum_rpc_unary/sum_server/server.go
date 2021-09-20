package main

import (
	"context"
	"fmt"
	"kubernetes-playground/udemy.grpc-golang/sum_rpc_unary/sumpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	sumpb.UnimplementedSumServiceServer
}

func (*server) Sum(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	result := req.GetSummand().GetFirstInteger() + req.GetSummand().GetSecondInteger()
	response := sumpb.SumResponse{
		Result: result,
	}
	return &response, nil
}

func main() {
	fmt.Println("Server starting...")
	lis, error := net.Listen("tcp", "0.0.0.0:5001")
	if error != nil {
		log.Fatalf("Server could not listen: %v", error)
	}
	s := grpc.NewServer()
	sumpb.RegisterSumServiceServer(s, &server{})

	err := s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
