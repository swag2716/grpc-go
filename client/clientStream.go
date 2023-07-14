package main

import (
	"context"
	"log"
	"time"

	pb "github.com/swapnika/grpc-go/proto"
)

func callSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Client streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())

	if err != nil {
		log.Fatal("count not send names : ", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatal("Error while sending : ", err)
		}

		log.Println("Send the request with name : ", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	log.Println("Client Streaming finished")

	if err != nil {
		log.Fatal("Error while recieving : ", err)
	}

	log.Println(res.Messages)

}
