package main

import (
	"context"
	"log"
	"time"

	pb "github.com/caspereijkens/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, duration time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), duration)

	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Casper",
	}

	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded !")
				return
			} else {
				log.Fatalf("Unexpected gRPC error: %v\n", e)
			}
		} else {
			log.Fatalf("A non-gRPC error: %v\n", err)
		}
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
