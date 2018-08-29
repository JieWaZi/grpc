package server

import (
	"context"
	pb "github.com/grpc/proto"
	"net"
	"log"
	"google.golang.org/grpc"
	"time"
	"github.com/grpc/api/grpc-gateway/http/handler"
	"github.com/grpc/api/consul"
)

type SecondCalculationService struct{}

func (s *SecondCalculationService) SubCalculation(ctx context.Context, request *pb.SubCalculationRequest) (*pb.SubCalculationResponse, error) {
	return &pb.SubCalculationResponse{
		Result: request.First - request.Second,
	}, nil
}

func (*SecondCalculationService) Start() {
	lis, err := net.Listen("tcp", "127.0.0.1:9200")
	if err != nil {
		log.Fatal("SecondCalculationService: " + err.Error())
	}
	s := grpc.NewServer()
	pb.RegisterSecondCalculationServiceServer(s, &SecondCalculationService{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("SecondCalculationService: " + err.Error())
	}
}

func init() {
	consul.Register("SecondCalculationService", "127.0.0.1", 9200, "127.0.0.1:8500", 10*time.Second, 15)
	handler.RegisterHTTPHandler("SecondCalculationService", pb.RegisterSecondCalculationServiceHandler)
}
