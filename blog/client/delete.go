package main

import (
	"context"
	"log"

	pb "github.com/caspereijkens/grpc-go-course/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("---deleteBlog was invoked---")

	req := &pb.BlogId{
		Id: id,
	}

	_, err := c.DeleteBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("Unexpected error %v\n", err)
	}

	log.Println("Blog has been deleted")
}
