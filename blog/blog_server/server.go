package main

import (
	"context"
	"fmt"
	blogpb "github.com/SarathLUN/udemy-grpc-go-course/blog/blog_pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
)

type server struct {
	blogpb.UnimplementedBlogServiceServer
}

func (s server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	blog := req.GetBlog()
	log.Printf("creating blog request: %v", blog)
	data := blogItem{
		AuthorID: blog.GetAuthor(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}
	// insert
	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	// get inserted ID
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("cannot convert to ObjectID: %v", err))
	}
	// return the CreateBlogResponse
	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:      oid.Hex(),
			Author:  blog.GetAuthor(),
			Title:   blog.GetTitle(),
			Content: blog.GetContent(),
		},
	}, nil
}

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

var collection *mongo.Collection

func main() {
	// if we crush the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// connect to MongoDB
	log.Println("Connecting to MongoDB")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("failed to create mongo client: %v", err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatalf("failed to connect mongoDB: %v", err)
	}
	// create collection
	collection = client.Database("mydb").Collection("blog")

	// starting messages
	log.Println("Blog Server started...")
	log.Println("Ctrl+C to stop the server")
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, server{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// wait for Ctrl+C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// block until a signal is received
	<-ch
	// stop the server
	log.Println("Stopping the server")
	s.Stop()
	// close the listener
	log.Println("Closing the listener")
	_ = lis.Close()
	// close connection with MongoDB
	log.Println("Closing connection with MongoDB")
	_ = client.Disconnect(context.TODO())
	// End of program
	log.Println("End of Program")

}