package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/swapnika/grpc-go/proto"
)

func callSayHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Bidirectional streaming has started")

	stream, err := client.SayHelloBidirectionalStreaming(context.Background())

	if err != nil {
		log.Fatal("count not send names : ", err)
	}

	waitc := make(chan struct{})

	go func() {
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
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatal("Error while sending request : ", err)
		}
		time.Sleep(2 * time.Second)
	}

	stream.CloseSend()
	<-waitc
	log.Println("Bidirectional streaming finished")
}
