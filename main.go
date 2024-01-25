package main

import (
	"cosgrpctest/protos"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"cosgrpctest/internal/server"
)

func main() {
	s := grpc.NewServer()

	reflection.Register(s)

	protos.RegisterCosServiceServer(s, &server.CosServer{})

	lis, err := net.Listen("tcp", ":9111")
	if err != nil {
		panic(err)
	}

	fmt.Println("Server listening on port 9111")

	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
