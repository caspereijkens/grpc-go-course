package main

import (
	"context"
	"log"

	pb "github.com/caspereijkens/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  25,
		SecondNumber: -55,
	})

	if err != nil {
		log.Fatalf("Could not calculate: %v\n", err)
	}

	log.Printf("Sum is: %d\n", res.Result)
}
