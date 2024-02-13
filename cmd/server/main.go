package main

import (
	"log"
	"net"
	"test/pkg/adder"
	"test/pkg/api/pkg/api"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	src := &adder.GRPCServer{}
	api.RegisterNumberServer(s, src)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
