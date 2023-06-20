package db

import (
	"context"
	match_frontend "match_frontend/proto"
)

var Default Service

type Service interface {
	Init(ctx context.Context, opts ...Option) error
	Close(ctx context.Context) error
	String() string
	AddToken(ctx context.Context, info *match_frontend.MatchInfo) error
	RemoveToken(ctx context.Context, matchId string, gameId string, subType int64) error
}

type MatchInfo struct {
	Id       string
	PlayerId string
	Score    int64
	GameId   string
	SubType  int64
}
