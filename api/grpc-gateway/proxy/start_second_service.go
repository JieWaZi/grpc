package main

import "github.com/grpc/api/grpc-gateway/server"

func main() {

	secondCalculationService := server.SecondCalculationService{}
	secondCalculationService.Start()
}
