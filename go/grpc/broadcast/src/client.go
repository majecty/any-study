package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "majecty.com/broadcast/protocol"
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
	client := pb.NewBroadcastServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	sendToAll(client, ctx)
	receiveAll(client, ctx)
}

func receiveAll(
	client pb.BroadcastServiceClient, ctx context.Context) {
	stream, err := client.MessagesFromOthers(
		ctx,
		&pb.Empty{Message: ""},
	)
	if err != nil {
		log.Fatalf("failed to get messages from others, %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break;
		}
		if err != nil {
			log.Fatalf("failed to get a message from others, %v", err)
		}
		log.Printf("msg from others: %v", msg);
	}
}

func sendToAll(client pb.BroadcastServiceClient, ctx context.Context) {
	hello, err := client.SendToAll(ctx, &pb.BroadcastMessage{
		From:      "a",
		To:        "b",
		Message:   "c",
		Broadcast: true,
	})
	if err != nil {
		log.Fatalf("failed say hello %v", err)
	}
	log.Printf("hello %v", hello)
}
