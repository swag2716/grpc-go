package main

import (
	"context"
	"log"
	"time"

	pb "github.com/swapnika/grpc-go/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})

	if err != nil {
		log.Fatal("Could not greet", err)
	}

	log.Fatal(res.Message)
}
