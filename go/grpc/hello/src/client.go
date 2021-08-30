package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "majecty.com/hello/protocol"
	"time"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	hello, err := client.SayHello(ctx, &pb.HelloRequest{Name: "hi"})
	if err != nil {
		log.Fatalf("failed say hello %v", err);
	}
	log.Printf("hello %v", hello)
}
