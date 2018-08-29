package main

// server.go

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/grpc/proto"
)

type grpcServer struct{}

func (s *grpcServer) SendMessage(ctx context.Context, in *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	return &pb.SendMessageResponse{FromWho: in.ToWho, Message: "receive message : " + in.Message}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMessageServer(s, &grpcServer{})
	s.Serve(lis)
}
