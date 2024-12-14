package main

import (
	"log"

	pb "github.com/caspereijkens/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) Factorize(in *pb.PrimeFactorRequest, stream grpc.ServerStreamingServer[pb.PrimeFactorResponse]) error {
	log.Printf("Factorize function was invoked with %v\n", in)

	N := int(in.Integer)

	p := 2

	for N > 1 {
		if (N % p) > 0 {
			p += 1
			continue
		}

		stream.Send(&pb.PrimeFactorResponse{
			PrimeFactor: int32(p),
		})

		N = N / p
	}
	return nil
}
