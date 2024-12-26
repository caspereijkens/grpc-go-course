package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/caspereijkens/grpc-go-course/blog/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListBlogs(in *empty.Empty, stream grpc.ServerStreamingServer[pb.Blog]) error {
	log.Println("ListBlogs was invoked")

	cursor, err := collection.Find(context.Background(), primitive.D{{}})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot list blogs in collection: %v\n", err),
		)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		data := &BlogItem{}
		if err := cursor.Decode(data); err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error decoding document from MongoDB: %v\n", err),
			)
		}

		err = stream.Send(documentToBlog(data))
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error sending blog on stream: %v\n", err),
			)
		}
	}

	if err = cursor.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v\n", err),
		)
	}

	return nil
}
