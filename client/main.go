package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/n6o/grpc-go-sandbox/gen-client"
)

const (
	defaultVoice = "Heeeey"
)

var (
	addr  = flag.String("addr", "localhost:50051", "the address to connect to")
	voice = flag.String("voice", defaultVoice, "Name to greet")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Echo(ctx, &pb.EchoRequest{Soound: *voice})
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}
	log.Printf("Echo: %s", r.GetSoound())
}
