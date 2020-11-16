package main

import (
	"context"
	"fmt"
	"net"

	"github.com/Hexnaught/lrn-grpc-redis/database"
	"github.com/Hexnaught/lrn-grpc-redis/proto"
	"google.golang.org/grpc"
)

type server struct {
	// We have to implement this in our server struct to satisfy
	// the protobuf auto-generated interfaces
	proto.UnimplementedBasicServiceServer
}

var db database.Database

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	// We can take this as an env var to switch DB being used
	// and our basic level of abstraction allows us to easily plug in
	// mongo or similar
	db, err = database.Factory("redis")
	if err != nil {
		panic(err)
	}

	proto.RegisterBasicServiceServer(srv, &server{})

	fmt.Printf("Starting server on :4040\n")

	if err = srv.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *server) Set(ctx context.Context, in *proto.SetRequest) (*proto.ServerResponse, error) {
	value, err := db.Set(in.GetKey(), in.GetValue())
	return generateResponse(value, err)
}

func (s *server) Get(ctx context.Context, in *proto.GetRequest) (*proto.ServerResponse, error) {
	value, err := db.Get(in.GetKey())
	return generateResponse(value, err)
}

func (s *server) Delete(ctx context.Context, in *proto.DeleteRequest) (*proto.ServerResponse, error) {
	value, err := db.Delete(in.GetKey())
	return generateResponse(value, err)
}

func generateResponse(value string, err error) (*proto.ServerResponse, error) {
	if err != nil {
		return &proto.ServerResponse{Success: false, Value: value, Error: err.Error()}, nil
	}
	return &proto.ServerResponse{Success: true, Value: value, Error: ""}, nil
}
