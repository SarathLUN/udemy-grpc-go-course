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
	// create blog
	//doCreateBlog(c)
	// read blog
	//doReadBlog(c)
	// update blog
	doUpdateBlog(c)
}

func doUpdateBlog(c blogpb.BlogServiceClient) {
	newBlog := &blogpb.Blog{
		Id:      "615946967829cff00b77ff5f",
		Author:  "Changed Author",
		Title:   "My First Blog (edited)",
		Content: "content of the first blog, with some awesome additions!",
	}
	ub, err := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: newBlog})
	if err != nil {
		log.Fatalf("client: failed to update blog: %v", err)
	}
	log.Printf("client: updated blog: %v", ub)

}

func doReadBlog(c blogpb.BlogServiceClient) {
	log.Println("Client: doReadBlog")
	res, err := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "21594d7bcd817d3d629f5138"})
	if err != nil {
		log.Fatalf("client: failed to read blog: %v", err)
	}
	log.Printf("client: read blog: %v", res)
}

func doCreateBlog(c blogpb.BlogServiceClient) {
	log.Println("Client: doCreateBlog")
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
