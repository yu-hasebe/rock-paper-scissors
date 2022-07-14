package main

import (
	"fmt"
	"log"

	"github.com/yu-hasebe/rock-paper-scissors/pb"
	"github.com/yu-hasebe/rock-paper-scissors/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := 50051
	listenPort, err := net.listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterRockPaperScissorsServiceServer(server, service.NewRockPaperScissorsService())

	reflection.Register(server)
	server.Serve(listenPort)
}
