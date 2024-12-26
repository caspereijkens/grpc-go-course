package main

import (
	"context"
	"io"
	"log"

	pb "github.com/caspereijkens/grpc-go-course/blog/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("---listBlog was invoked---")

	stream, err := c.ListBlogs(context.Background(), &empty.Empty{})

	if err != nil {
		log.Fatalf("Error while calling ListBlogs: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while receiving from stream: %v\n", err)
		}

		log.Println(res)
	}
}
