package main

import "github.com/grpc/api/grpc-gateway/server"

func main() {
	firstCalculationService := server.FirstCalculationService{}
	firstCalculationService.Start()
}