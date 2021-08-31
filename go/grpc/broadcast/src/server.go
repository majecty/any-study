package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "majecty.com/broadcast/protocol"
	"net"
	"sync"
)

const (
	port = ":50051"
)

type Client struct {
	toClient   chan *pb.BroadcastMessage
	fromClient chan string
}

func NewClient() *Client {
	return &Client{
		toClient:   make(chan *pb.BroadcastMessage),
		fromClient: make(chan string),
	}
}

func (client Client) DisconnectFromClient() {
	close(client.fromClient)
}

type server struct {
	pb.UnimplementedBroadcastServiceServer
	msgChan chan pb.BroadcastMessage
	clients []*Client
	// used to modify clients
	mu sync.Mutex
}

func NewServer() *server {
	return &server{
		msgChan: make(chan pb.BroadcastMessage),
	}
}

func (s *server) SendToAll(ctx context.Context, msg *pb.BroadcastMessage) (*pb.SendToAllResponse, error) {
	log.Printf("receive msg %v", msg)

	// we need to lock because s.clients can be changed
	s.mu.Lock()
	defer s.mu.Unlock()

	clients := s.clients
	log.Printf("send message %v to clients", msg)
	for _, client := range clients {
		log.Print(".")
		client.toClient <- msg
	}

	return &pb.SendToAllResponse{Message: "received well"}, nil
}

func (s *server) AddClient(client *Client) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clients = append(s.clients, client)
	log.Printf("add client. len(clients) = %v", len(s.clients))
}

func (s *server) RemoveClient(client *Client) {
	s.mu.Lock()
	defer s.mu.Unlock()

	clientIndex := 0
	for i, c := range s.clients {
		if c == client {
			clientIndex = i
			break
		}
	}

	s.clients = append(s.clients[:clientIndex], s.clients[clientIndex+1:]...)

	log.Printf("remove client. len(clients) = %v", len(s.clients))
}

func (s *server) MessagesFromOthers(_ *pb.Empty, stream pb.BroadcastService_MessagesFromOthersServer) error {
	client := NewClient()
	defer client.DisconnectFromClient()

	s.AddClient(client)
	defer s.RemoveClient(client)

	for message := range client.toClient {
		err := stream.Send(message)
		if err != nil {
			log.Printf("failed to send message to client %v", err)
			break
		}
	}
	log.Printf("toClient is closed")
	return nil
}

func main() {
	fmt.Println("Hello, World!")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBroadcastServiceServer(s, NewServer())
	log.Printf("server listening at %v", lis.Addr())

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
