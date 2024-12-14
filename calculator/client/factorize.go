package main

import (
	"context"
	"io"
	"log"

	pb "github.com/caspereijkens/grpc-go-course/calculator/proto"
)

func doFactorization(c pb.CalculatorServiceClient, number uint64) {
	log.Println("doFactorization was invoked")

	req := &pb.PrimeFactorRequest{
		Number: number,
	}

	stream, err := c.Factorize(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}
		log.Printf("Prime Factor: %d\n", msg.Result)
	}
}
