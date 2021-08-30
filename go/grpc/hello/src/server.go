package main

// golang import from local
import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "majecty.com/hello/protocol"
	"net"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{
		Message: "Hello " + in.GetName(),
	}, nil
}

func main() {
	fmt.Println("Hello, World!")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
