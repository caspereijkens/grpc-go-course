package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/caspereijkens/grpc-go-course/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	numbers := []int64{3, 5, 9, 54, 23, 8, 13, 21, 34, 55, 89}

	waitc := make(chan struct{})

	go func() {
		for _, number := range numbers {
			log.Printf("Send request: %d\n", number)
			stream.Send(&pb.MaxRequest{Number: number})
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving response from stream: %v\n", err)
				break
			}

			log.Printf("Received: %d\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
