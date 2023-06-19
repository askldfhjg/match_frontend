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

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Match_frontend) Stream(ctx context.Context, req *match_frontend.StreamingRequest, stream match_frontend.Match_frontend_StreamStream) error {
	log.Infof("Received Match_frontend.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&match_frontend.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Match_frontend) PingPong(ctx context.Context, stream match_frontend.Match_frontend_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&match_frontend.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
