package main

import (
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/caspereijkens/grpc-go-course/blog/proto"
)

var addr string = "localhost:50051"

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id)
	// readBlog(c, "aNonExistingId")
	updateBlog(c, id)
	listBlog(c)
	deleteBlog(c, id)
}
