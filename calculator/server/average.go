package main

import (
	"io"
	"log"

	pb "github.com/caspereijkens/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) Avg(stream grpc.ClientStreamingServer[pb.AvgRequest, pb.AvgResponse]) error {
	var avg float64
	var sum int32
	var count int32

	log.Printf("Avg function was invoked with %v\n", stream)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: avg,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		log.Printf("Receiving request: %v\n", req)
		count++
		sum += req.Number
		avg = float64(sum) / float64(count)
	}
}
