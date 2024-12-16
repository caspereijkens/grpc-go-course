package main

import (
	"context"
	"log"
	"time"

	pb "github.com/caspereijkens/grpc-go-course/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")

	numbers := []int32{3, 5, 9, 54, 23}

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Avg: %v\n", err)
	}

	for _, number := range numbers {
		log.Printf("Sending number: %v\n", number)
		err = stream.Send(&pb.AvgRequest{Number: number})
		time.Sleep(1 * time.Second)

		if err != nil {
			log.Fatalf("Error while sending the stream: %v\n", err)
		}
	}
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving the response from Avg: %v\n", err)
	}

	log.Printf("Avg: %.6f\n", res.Result)
}
