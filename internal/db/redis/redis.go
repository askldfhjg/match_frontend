package redis

import (
	"context"
	"fmt"
	match_frontend "match_frontend/proto"

	"github.com/gomodule/redigo/redis"
	"github.com/micro/micro/v3/service/logger"
	"google.golang.org/protobuf/proto"
)

const PlayerInfoExpireTime = 300

func (m *redisBackend) AddToken(ctx context.Context, info *match_frontend.MatchInfo) error {
	redisConn, err := m.redisPool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer handleConnectionClose(&redisConn)

	playerId := info.GetPlayerId()
	value, err := proto.Marshal(info)
	if err != nil {
		return err
	}
	if value == nil {
		return fmt.Errorf("failed to marshal the ticket proto, id: %s: proto: Marshal called with nil", playerId)
	}

	//remove old
	key := fmt.Sprintf(ticketKey, playerId)
	bb, err := redis.Bytes(redisConn.Do("GET", key))
	if err != nil && err != redis.ErrNil {
		return err
	}
	t := &match_frontend.MatchInfo{}
	err = proto.Unmarshal(bb, t)
	if err != nil {
		logger.Error("AddToken frist get %s have err %s", playerId, err.Error())
	} else {
		zsetKey := fmt.Sprintf(allTickets, t.GameId, t.SubType)
		redisConn.Do("ZREM", zsetKey, playerId)
	}

	//add new
	_, err = redisConn.Do("SET", key, value, "EX", PlayerInfoExpireTime)
	if err != nil {
		return err
	}
	zsetKey := fmt.Sprintf(allTickets, info.GameId, info.SubType)
	_, err = redisConn.Do("ZADD", zsetKey, info.Score, playerId)
	if err != nil {
		_, errs := redisConn.Do("DEL", key)
		if errs != nil {
			logger.Error("ZADD and DEL %s have err %s", playerId, errs.Error())
		}
		return err
	}
	return nil
}

func (m *redisBackend) RemoveToken(ctx context.Context, playerId string, gameId string, subType int64) error {
	redisConn, err := m.redisPool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer handleConnectionClose(&redisConn)

	key := fmt.Sprintf(ticketKey, playerId)
	bb, err := redis.Bytes(redisConn.Do("GET", key))
	if err != nil {
		if err != redis.ErrNil {
			return err
		} else {
			return nil
		}
	}
	t := &match_frontend.MatchInfo{}
	err = proto.Unmarshal(bb, t)
	del := false
	if err != nil {
		logger.Error("RemoveToken frist get %s have err %s", playerId, err.Error())
	} else {
		del = true
		zsetKey := fmt.Sprintf(allTickets, t.GameId, t.SubType)
		redisConn.Do("ZREM", zsetKey, playerId)

	}
	if t.GameId != gameId || t.SubType != subType || !del {
		zsetKey := fmt.Sprintf(allTickets, gameId, subType)
		redisConn.Do("ZREM", zsetKey, playerId)
	}

	_, err = redisConn.Do("DEL", key)
	if err != nil {
		logger.Error("RemoveToken %s have err %s", playerId, err.Error())
	}
	return nil

}

func (m *redisBackend) GetToken(ctx context.Context, playerId string) (*match_frontend.MatchInfo, error) {
	redisConn, err := m.redisPool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer handleConnectionClose(&redisConn)
	key := fmt.Sprintf(ticketKey, playerId)

	bb, err := redis.Bytes(redisConn.Do("GET", key))
	if err != nil {
		if err != redis.ErrNil {
			return nil, err
		} else {
			return &match_frontend.MatchInfo{}, nil
		}
	}
	t := &match_frontend.MatchInfo{}
	err = proto.Unmarshal(bb, t)
	if err != nil {
		return nil, err
	}
	zsetKey := fmt.Sprintf(allTickets, t.GameId, t.SubType)
	_, err = redis.Int(redisConn.Do("ZSCORE", zsetKey, playerId))
	if err != nil {
		if err != redis.ErrNil {
			return nil, err
		} else {
			redisConn.Do("DEL", key)
			return &match_frontend.MatchInfo{}, nil
		}
	}
	return t, nil
}
