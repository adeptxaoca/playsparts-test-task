package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "part_handler/internal/pkg/api/v1"
)

func createPart(client pb.PartServiceClient, part *pb.Part) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if res, err := client.Create(ctx, &pb.CreateReq{Part: part}); err != nil {
		log.Printf("ERROR CREATE: %+v\n", err)
	} else {
		log.Printf("CREATE: %v", res.Part)
	}
}

func readPart(client pb.PartServiceClient, id uint64) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if res, err := client.Read(ctx, &pb.ReadReq{Id: id}); err != nil {
		log.Printf("ERROR READ [%d]: %v\n", id, err)
	} else {
		log.Printf("READ: %v", res.Part)
	}
}

func main() {
	conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() { _ = conn.Close() }()

	// Creates a new CustomerClient
	client := pb.NewPartServiceClient(conn)

	// Test queries
	readPart(client, 1)
	readPart(client, 10)
	createPart(client, &pb.Part{ManufacturerId: 1, Name: "A&D", VendorCode: "#V1247-100#"})
	//createPart(client, &pb.Part{ManufacturerId: 100, Name: "A&D", VendorCode: "V1247-100####"})
	createPart(client, &pb.Part{ManufacturerId: 100, Name: "my first part", VendorCode: "V1247"})
	createPart(client, &pb.Part{ManufacturerId: 2, Name: "second part", VendorCode: "8QQAOXJZBV"})
	createPart(client, &pb.Part{ManufacturerId: 100, Name: "my first part", VendorCode: "8QQAOXJZBV"})
}
