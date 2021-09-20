package main

import (
	"context"
	"fmt"
	"log"

	"kubernetes-playground/udemy.grpc-golang/code_generation_test/greet/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello i'm a client")
	conn, err := grpc.Dial("localhost:5005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	// fmt.Printf("Created client: %f", c)

	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Simon",
			LastName:  "Schmid",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet rpc: %v", err)
	}
	log.Printf("Response from greet: %v", res.Result)
}
