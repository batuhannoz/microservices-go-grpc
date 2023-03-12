package client

import (
	"context"
	"github.com/batuhannoz/microservices-go-grpc/starwars/service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type StarWarsClient struct {
	client proto.StarWarsServiceClient
}

func NewStarWarsClient() *StarWarsClient {
	conn, err := grpc.Dial("localhost"+":3001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewStarWarsServiceClient(conn)
	return &StarWarsClient{
		client: client,
	}
}

func (s *StarWarsClient) GetRandomCharacter() *proto.Character {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := s.client.RandomCharacter(ctx, &proto.NoParam{})
	if err != nil {
		log.Fatal(err)
	}
	return res
}
