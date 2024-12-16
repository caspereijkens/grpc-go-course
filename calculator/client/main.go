package main

import (
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/caspereijkens/grpc-go-course/calculator/proto"
)

var addr string = "localhost:50052"

func main() {
	// number := flag.Int("number", 0, "The number to be used in factorization")
	flag.Parse()

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	// doFactorization(c, uint64(*number))
	numbers := []int32{3, -5, -9, 54, -23, 8, -13, 21, -34, 55, -89}

	for _, number := range numbers {
		doSqrt(c, number)
	}
}
