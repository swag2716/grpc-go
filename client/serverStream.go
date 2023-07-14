package main

import (
	"context"
	"io"
	"log"

	pb "github.com/swapnika/grpc-go/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)

	if err != nil {
		log.Fatal("could not send name: ", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("Error while streaming : ", err)
		}

		log.Println(message)
	}

	log.Println("Streaming finished")
}
