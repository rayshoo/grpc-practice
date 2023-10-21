package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "prac0/grpc/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:8080", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAuthServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Login(ctx, &pb.LoginRequest{Username: "admin",Password: "qwerty"})
	if err != nil {
		log.Fatalf("could not login: %v", err)
	}
	log.Printf("Login: %s", r.GetLoginCode())
	log.Printf("Token: %s", r.GetToken())
}