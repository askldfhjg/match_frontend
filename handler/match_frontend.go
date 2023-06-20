package handler

import (
	"context"
	"match_frontend/internal/db"
	match_frontend "match_frontend/proto"

	"github.com/micro/micro/v3/service/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type Handler struct{}

func (e *Handler) EnterMatch(ctx context.Context, req *match_frontend.EnterMatchReq, rsp *match_frontend.EnterMatchRsp) error {
	logger.Infof("Received EnterMatch request with count: %+v", req)

	matchInfo, ok := proto.Clone(req.Param).(*match_frontend.MatchInfo)
	if !ok {
		return status.Error(codes.Internal, "failed to clone input matchInfo proto")
	}

	playerId := matchInfo.GetPlayerId()
	if len(playerId) <= 0 {
		rsp.Code = -1
		rsp.Err = "playerId <= 0"
		return nil
	}
	gameId := matchInfo.GetGameId()
	subType := matchInfo.GetSubType()
	if len(gameId) <= 0 || subType <= 0 {
		rsp.Code = -1
		rsp.Err = "gameId err or subType err"
		return nil
	}
	if matchInfo.GetScore() <= 0 {
		rsp.Code = -1
		rsp.Err = "score error"
		return nil
	}
	err := db.Default.AddToken(ctx, matchInfo)
	if err != nil {
		logger.Errorf("EnterMatch AddToken err : %s", err.Error())
		rsp.Code = -1
		rsp.Err = err.Error()
		return nil
	}
	return nil
}

func (e *Handler) LevelMatch(ctx context.Context, req *match_frontend.LevelMatchReq, rsp *match_frontend.LevelMatchRsp) error {
	playerId := req.GetPlayerId()
	if len(playerId) <= 0 {
		rsp.Code = -1
		rsp.Err = "playerId <= 0"
		return nil
	}
	gameId := req.GetGameId()
	subType := req.GetSubType()
	if len(gameId) <= 0 || subType <= 0 {
		rsp.Code = -1
		rsp.Err = "gameId err or subType err"
		return nil
	}
	err := db.Default.RemoveToken(ctx, req.GetPlayerId(), req.GetGameId(), req.GetSubType())
	if err != nil {
		logger.Errorf("LevelMatch RemoveToken err : %s", err.Error())
		rsp.Code = -1
		rsp.Err = err.Error()
		return nil
	}
	return nil
}

func (e *Handler) GetMatchInfo(ctx context.Context, req *match_frontend.GetMatchInfoReq, rsp *match_frontend.GetMatchInfoRsp) error {
	playerId := req.GetPlayerId()
	if len(playerId) <= 0 {
		rsp.Code = -1
		rsp.Err = "playerId <= 0"
		return nil
	}
	info, err := db.Default.GetToken(ctx, req.GetPlayerId())
	if err != nil {
		logger.Errorf("LevelMatch AddToken err : %s", err.Error())
		rsp.Code = -1
		rsp.Err = err.Error()
		return nil
	}
	rsp.Result = info
	return nil
}
