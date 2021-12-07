package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	chat "grpc/chat"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	msg := chat.Message{
		Body: "Hello from the client!",
	}

	resp, err := c.SayHello(context.Background(), &msg)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from Server: %s", resp.Body)
}
