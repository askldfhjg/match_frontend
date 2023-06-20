package handler

import (
	"context"
	"fmt"
	"match_frontend/internal/db"
	match_frontend "match_frontend/proto"

	log "github.com/micro/micro/v3/service/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type Handler struct{}

func (e *Handler) EnterMatch(ctx context.Context, req *match_frontend.EnterMatchReq, rsp *match_frontend.EnterMatchRsp) error {
	log.Infof("Received EnterMatch request with count: %+v", req)

	matchInfo, ok := proto.Clone(req.Param).(*match_frontend.MatchInfo)
	if !ok {
		return status.Error(codes.Internal, "failed to clone input matchInfo proto")
	}

	matchInfo.Id = fmt.Sprintf("%s:%s:%d", matchInfo.GetPlayerId(), matchInfo.GetGameId(), matchInfo.GetSubType())
	err := db.Default.AddToken(ctx, matchInfo)
	if err != nil {
		log.Errorf("EnterMatch AddToken err : %s", err.Error())
		return err
	}
	rsp.MatchId = matchInfo.Id
	return nil
}

func (e *Handler) LevelMatch(ctx context.Context, req *match_frontend.LevelMatchReq, rsp *match_frontend.LevelMatchRsp) error {
	err := db.Default.RemoveToken(ctx, req.MatchId, req.GameId, req.SubType)
	if err != nil {
		log.Errorf("LevelMatch AddToken err : %s", err.Error())
		return err
	}
	rsp.Err = ""
	return nil
}
