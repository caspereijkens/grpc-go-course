package main

import (
	"context"
	"log"
	"time"

	pb "github.com/caspereijkens/grpc-go-course/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Casper"},
		{FirstName: "Yuanyuan"},
		{FirstName: "Pegi"},
		{FirstName: "Noma"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		err = stream.Send(req)
		time.Sleep(1 * time.Second)

		if err != nil {
			log.Fatalf("Error while sending the stream: %v\n", err)
		}
	}
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving the response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
