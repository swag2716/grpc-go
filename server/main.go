package main

import (
	"log"
	"net"

	pb "github.com/swapnika/grpc-go/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal("Failed to start the server", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Println("Server started at", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to start", err)
	}

}
