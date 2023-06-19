package main

import (
	"match_frontend/handler"
	pb "match_frontend/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("match_frontend"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterMatchFrontendHandler(srv.Server(), new(handler.Match_frontend))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
