package main

//client.go

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/grpc/proto"
)
func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMessageClient(conn)

	r, err := c.SendMessage(context.Background(), &pb.SendMessageRequest{ToWho: "wjj",Message:"helloWorld"})
	if err != nil {
		log.Fatal("could not SendMessage: %v", err)
	}
	log.Printf("SendMessage: %s", r.Message)
}