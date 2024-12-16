package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/caspereijkens/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) LongGreet(stream grpc.ClientStreamingServer[pb.GreetRequest, pb.GreetResponse]) error {
	log.Printf("LongGreet function was invoked with %v\n", stream)

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		log.Printf("Receiving request: %v\n", req)

		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}
