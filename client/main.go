package main

import (
	"log"

	pb "github.com/swapnika/grpc-go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Did not connect", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Swapnika", "Alice", "Bob"},
	}

	// callSayHello(client)
	callSayHelloServerStreaming(client, names)
}
