package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	match_frontend "match_frontend/proto"
)

type Match_frontend struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Match_frontend) Call(ctx context.Context, req *match_frontend.Request, rsp *match_frontend.Response) error {
	log.Info("Received Match_frontend.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
