package server

import (
	"context"
	pb "github.com/grpc/proto"
	"net"
	"log"
	"google.golang.org/grpc"
	"time"
	"github.com/grpc/api/consul"
	"github.com/grpc/api/grpc-gateway/http/handler"
)

type FirstCalculationService struct{}

func (s *FirstCalculationService) AddCalculation(ctx context.Context, request *pb.AddCalculationRequest) (*pb.AddCalculationResponse, error) {
	return &pb.AddCalculationResponse{
		Result: request.First + request.Second,
	}, nil
}

func (*FirstCalculationService) Start() {
	lis, err := net.Listen("tcp", "127.0.0.1:9100")
	if err != nil {
		log.Fatal("FirstCalculationService: " + err.Error())
	}
	s := grpc.NewServer()
	pb.RegisterFirstCalculationServiceServer(s, &FirstCalculationService{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("FirstCalculationService: " + err.Error())
	}
}

func init() {
	consul.Register("FirstCalculationService", "127.0.0.1", 9100, "127.0.0.1:8500", 10*time.Second, 15)
	handler.RegisterHTTPHandler("FirstCalculationService", pb.RegisterFirstCalculationServiceHandler)
}
