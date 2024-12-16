package main

import (
	"io"
	"log"

	pb "github.com/caspereijkens/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) Max(stream grpc.BidiStreamingServer[pb.MaxRequest, pb.MaxResponse]) error {
	var max int64
	var set bool

	log.Printf("Max function was invoked with %v\n", stream)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		log.Printf("Receiving request: %v\n", req)

		if req.Number > max || !set {
			set = true
			max = req.Number
		}

		err = stream.Send(&pb.MaxResponse{
			Result: max,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}
