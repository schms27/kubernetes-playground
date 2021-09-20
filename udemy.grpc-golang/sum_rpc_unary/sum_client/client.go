package main

import (
	"context"
	"kubernetes-playground/udemy.grpc-golang/sum_rpc_unary/sumpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client could not connect: %v", err)
	}
	defer conn.Close()

	client := sumpb.NewSumServiceClient(conn)

	req := &sumpb.SumRequest{
		Summand: &sumpb.Summand{
			FirstInteger:  12,
			SecondInteger: 24,
		},
	}
	res, err := client.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling sum rpc: %v", err)
	}
	log.Printf("Response from sum service: %v", res.Result)
}
