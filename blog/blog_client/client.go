package main

import (
	"context"
	blogpb "github.com/SarathLUN/udemy-grpc-go-course/blog/blog_pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	log.Println("Running blog client")
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Printf("failed to dial to server: %v", err)
	}
	defer func(cc *grpc.ClientConn) {
		_ = cc.Close()
	}(cc)
	c := blogpb.NewBlogServiceClient(cc)
	blog := &blogpb.Blog{
		Author:  "Tony",
		Title:   "My Second Blog",
		Content: "My Second blog content",
	}
	b, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("failed to create blog: %v", err)
	}
	log.Printf("blog has been created: %v", b)
}
