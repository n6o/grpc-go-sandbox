package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/n6o/grpc-go-sandbox/gen-client"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedEchoServer
}

func (s *server) SayEcho(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("Received: %v", in.GetSoound())
	return &pb.EchoResponse{Soound: "Echo " + in.GetSoound()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
