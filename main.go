package main

import (
	"log"
	"net"

	"protobuf/ganerete-file/github.com/protocolbuffers/protobuf/examples/go/tutorialpb"

	"protobuf/servers"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	tutorialpb.RegisterBookServer(grpcServer, &servers.BookServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
