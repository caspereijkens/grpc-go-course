package main

import (
	"context"
	"log"

	pb "github.com/caspereijkens/grpc-go-course/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---updateBlog was invoked---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "NotCasper",
		Title:    "New Title",
		Content:  "Content of the first blog with some awesome additions.",
	}
	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog has been updated!")
}
