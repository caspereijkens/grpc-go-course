package main

import (
	"log"
	"net"

	pb "github.com/caspereijkens/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on '%s': %v\n", addr, err)
	}

	log.Printf("Listening on %s\n", addr)

	opts := []grpc.ServerOption{}
	tls := true

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})
	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
