package main

import (
	"context"
	"github.com/batuhannoz/microservices-go-grpc/starwars/service/db"
	"github.com/batuhannoz/microservices-go-grpc/starwars/service/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":3001"
)

type StarWarsServer struct {
	proto.StarWarsServiceServer
}

func (s *StarWarsServer) RandomCharacter(ctx context.Context, req *proto.NoParam) (*proto.Character, error) {
	c := db.GetRandomCharacter()
	return &proto.Character{
		Name:      c.Name,
		Homeworld: c.Homeworld,
		ImageURL:  c.Pic,
	}, nil
}

func main() {
	db.InitializeDB()
	//listen on the port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
	// create a new gRPC server
	grpcServer := grpc.NewServer()
	// register StarWars service
	proto.RegisterStarWarsServiceServer(grpcServer, &StarWarsServer{})
	log.Printf("Server started at %v", lis.Addr())
	//list is the port, the grpc server needs to start there
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
