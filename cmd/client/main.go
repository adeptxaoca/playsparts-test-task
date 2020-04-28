package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "part_handler/internal/pkg/api/v1"
)

func createPart(client pb.PartServiceClient, part *pb.Part) *pb.CreateRes {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.Create(ctx, &pb.CreateReq{Part: part})
	if err != nil {
		log.Printf("ERROR CREATE: %+v\n", err)
	} else {
		log.Printf("CREATE: %v", res.Part)
	}
	return res
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

func updatePart(client pb.PartServiceClient, part *pb.Part) *pb.UpdateRes {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.Update(ctx, &pb.UpdateReq{Part: part})
	if err != nil {
		log.Printf("ERROR UPDATE [%d]: %v\n", part.Id, err)
	} else {
		log.Printf("UPDATE: %v", res.Part)
	}
	return res
}

func deletePart(client pb.PartServiceClient, id uint64) *pb.DeleteRes {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.Delete(ctx, &pb.DeleteReq{Id: id})
	if err != nil {
		log.Printf("ERROR DELETE [%d]: %v\n", id, err)
	} else {
		log.Printf("UPDATE: %v", res)
	}
	return res
}

func main() {
	conn, err := grpc.Dial("127.0.0.1:10000", grpc.WithInsecure())
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
	createPart(client, &pb.Part{ManufacturerId: 100, Name: "my first part", VendorCode: "V1247"})
	createPart(client, &pb.Part{ManufacturerId: 2, Name: "second part", VendorCode: "8QQAOXJZBV"})
	createPart(client, &pb.Part{})
	updatePart(client, &pb.Part{Id: 100, Name: "first#", VendorCode: "8QQA$OXJ#ZBV"})
	updatePart(client, &pb.Part{Id: 8, Name: "simple name", VendorCode: "3000BGF"})

	p := createPart(client, &pb.Part{ManufacturerId: 5, Name: "engine 3000", VendorCode: "QWERTY"})
	if p == nil {
		return
	}

	readPart(client, p.Part.Id)
	updatePart(client, &pb.Part{Id: p.Part.Id, Name: "engine 4000"})
	deletePart(client, p.Part.Id)
	updatePart(client, &pb.Part{Id: p.Part.Id, Name: "engine 5000"})
}
