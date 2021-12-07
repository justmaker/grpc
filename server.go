package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	chat "grpc/chat"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := chat.Server{}

	srv := grpc.NewServer()

	chat.RegisterChatServiceServer(srv, &s)

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}

}
