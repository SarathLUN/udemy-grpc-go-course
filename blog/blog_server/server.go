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
	"gopkg.in/mgo.v2/bson"
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

func (s server) ReadBlog(ctx context.Context, req *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	log.Println("ReadBlog request...")
	blogID := req.GetBlogId()
	// get object id
	oid, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Server: cannot parse ID"),
		)
	}
	// create empty struct
	data := &blogItem{}
	// prepare filter
	filter := bson.M{"_id": oid}
	// query data
	res := collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Server: cannot find blog with specified ID: %v", err),
		)
	}
	return &blogpb.ReadBlogResponse{Blog: dataToBlogPb(data)}, nil
}

func dataToBlogPb(data *blogItem) *blogpb.Blog {
	return &blogpb.Blog{
		Id:      data.ID.Hex(),
		Title:   data.Title,
		Author:  data.AuthorID,
		Content: data.Content,
	}
}

func (s server) UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {
	log.Println("Update blog")
	b := req.GetBlog()
	oid, err := primitive.ObjectIDFromHex(b.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("server: cannot parse ID"))
	}
	//create empty struct
	data := &blogItem{}
	filter := bson.M{"_id": oid}
	// first we check is object existed
	fb := collection.FindOne(ctx, filter)
	if err := fb.Decode(data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("server: cannot found blog with specified ID: %v", err))
	}
	// we update our internal struct
	data.AuthorID = b.GetAuthor()
	data.Title = b.GetTitle()
	data.Content = b.GetContent()
	// update collection
	_, err = collection.ReplaceOne(ctx, filter, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("server: cannot update object in MongoDB: %v", err))
	}
	// return UpdateBlogResponse
	return &blogpb.UpdateBlogResponse{Blog: dataToBlogPb(data)}, nil
}

func (s server) DeleteBlog(ctx context.Context, req *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {
	log.Println("Delete blog")
	oid, err := primitive.ObjectIDFromHex(req.GetBlogId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("server: cannot parse ID"))
	}
	filter := bson.M{"_id": oid}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("server: cannot delete object in MongoDB: %v", err))
	}
	if res.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("server: cannot found object to delete: %v", err))
	}
	return &blogpb.DeleteBlogResponse{BlogId: req.GetBlogId()}, nil
}

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
