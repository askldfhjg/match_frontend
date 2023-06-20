package main

import (
	"match_frontend/handler"
	"match_frontend/internal/db"
	"match_frontend/internal/db/redis"
	pb "match_frontend/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("match_frontend"),
		service.Version("latest"),
		service.BeforeStart(func() error {
			svr, err := redis.New(
				db.WithAddress("127.0.0.1:6379"),
				db.WithPoolMaxActive(5),
				db.WithPoolMaxIdle(100),
				db.WithPoolIdleTimeout(300))
			if err != nil {
				return err
			}
			db.Default = svr
			return nil
		}),
	)

	// Register handler
	pb.RegisterMatchFrontendHandler(srv.Server(), new(handler.Handler))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
