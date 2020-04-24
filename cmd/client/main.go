package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "part_handler/pkg/api/v1"
)

func createPart(client pb.PartServiceClient, part *pb.Part) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.Create(ctx, &pb.CreateReq{Part: part})
	if err != nil {
		log.Fatalf("Could not create Part: %v", err)
	}
	log.Printf("A new Customer has been added with id: %v", resp.Part)
}

func readPart(client pb.PartServiceClient, id uint64) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := client.Read(ctx, &pb.ReadReq{Id: id})
	if err != nil {
		log.Fatalf("Could not read Part[%d]: %v", id, err)
	}
	log.Println(res)
}

func main() {
	conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() { _ = conn.Close() }()

	// Creates a new CustomerClient
	client := pb.NewPartServiceClient(conn)
	// createPart(client, &pb.Part{})
	readPart(client, 1)
	readPart(client, 2)
	readPart(client, 3)
	readPart(client, 10)
}
