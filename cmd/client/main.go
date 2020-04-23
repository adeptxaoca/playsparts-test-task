package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "part_handler/pkg/api/v1"
)

func createPart(client pb.PartServiceClient, part *pb.Part) {
	resp, err := client.Create(context.Background(), &pb.CreateReq{Part: part})
	if err != nil {
		log.Fatalf("Could not create Part: %v", err)
	}
	log.Printf("A new Customer has been added with id: %v", resp.Part)
}

func main() {
	conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Creates a new CustomerClient
	client := pb.NewPartServiceClient(conn)
	createPart(client, &pb.Part{})
}
