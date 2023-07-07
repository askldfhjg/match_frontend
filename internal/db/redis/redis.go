package redis

import (
	"context"
	"fmt"
	match_frontend "match_frontend/proto"
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"
	"github.com/micro/micro/v3/service/logger"
)

const (
	allTickets         = "allTickets:%v:%s"
	ticketKey          = "ticket:"
	poolVersionKey     = "poolVersionKey:"
	lastPoolVersionKey = "lastPoolVersionKey:%s:%d"
)

const PlayerInfoExpireTime = 600

func (m *redisBackend) AddToken(ctx context.Context, info *match_frontend.MatchInfo) error {
	redisConn, err := m.redisPool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer handleConnectionClose(&redisConn)

	playerId := info.GetPlayerId()

	//remove old
	key := ticketKey + playerId
	matchInfo, err := redis.String(redisConn.Do("GET", key))
	if err != nil && err != redis.ErrNil {
		return err
	}
	if len(matchInfo) > 0 {
		version, err := redis.Int64(redisConn.Do("HGET", poolVersionKey+matchInfo, "poolVersionKey"))
		if err != nil {
			return err
		}
		redisConn.Do("ZREM", fmt.Sprintf(allTickets, version, matchInfo), playerId)
	}

	//add new
	value := fmt.Sprintf("%s:%d", info.GameId, info.SubType)
	version, err := redis.Int64(redisConn.Do("HGET", poolVersionKey+value, "poolVersionKey"))
	if err != nil {
		return err
	}
	_, err = redisConn.Do("SET", key, value, "EX", PlayerInfoExpireTime)
	if err != nil {
		return err
	}
	zsetKey := fmt.Sprintf(allTickets, version, value)
	_, err = redisConn.Do("ZADD", zsetKey, info.Score, playerId)
	if err != nil {
		_, errs := redisConn.Do("DEL", key)
		if errs != nil {
			logger.Error("ZADD and DEL %s have err %s", playerId, errs.Error())
		}
		return err
	} else {
		nowVersion, err := redis.Int64(redisConn.Do("HGET", poolVersionKey+value, "poolVersionKey"))
		if err != nil {
			return err
		}
		if nowVersion != version {
			zsetKey := fmt.Sprintf(allTickets, nowVersion, value)
			redisConn.Do("ZADD", zsetKey, info.Score, playerId)
		}
	}
	return nil
}

func (m *redisBackend) RemoveToken(ctx context.Context, playerId string, gameId string, subType int64) error {
	redisConn, err := m.redisPool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer handleConnectionClose(&redisConn)

	key := ticketKey + playerId
	matchInfo, err := redis.String(redisConn.Do("GET", key))
	if err != nil && err != redis.ErrNil {
		return err
	}
	if len(matchInfo) > 0 {
		version, err := redis.Int64(redisConn.Do("HGET", poolVersionKey+matchInfo, "poolVersionKey"))
		if err != nil {
			return err
		}
		redisConn.Do("ZREM", fmt.Sprintf(allTickets, version, matchInfo), playerId)
	}

	value := fmt.Sprintf("%s:%d", gameId, subType)
	if matchInfo != value {
		version, err := redis.Int64(redisConn.Do("HGET", poolVersionKey+value, "poolVersionKey"))
		if err != nil {
			return err
		}
		zsetKey := fmt.Sprintf(allTickets, version, value)
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
	key := ticketKey + playerId

	matchInfo, err := redis.String(redisConn.Do("GET", key))
	if err != nil {
		if err != redis.ErrNil {
			return nil, err
		} else {
			return &match_frontend.MatchInfo{}, nil
		}
	}
	tmps := strings.Split(matchInfo, ":")
	if len(tmps) != 2 {
		redisConn.Do("DEL", key)
		return &match_frontend.MatchInfo{}, nil
	}
	gameId := tmps[0]
	subType, err := strconv.ParseInt(tmps[1], 10, 64)
	if err != nil {
		logger.Errorf("GetToken ParseInt %s have error %s", matchInfo, err.Error())
		redisConn.Do("DEL", key)
		return &match_frontend.MatchInfo{}, nil
	}

	versionMap, err := redis.StringMap(redisConn.Do("HGET", poolVersionKey+key))
	if err != nil && err != redis.ErrNil {
		return nil, err
	}
	oldzsetKey := fmt.Sprintf(allTickets, versionMap["lastPoolVersionKey"], matchInfo)
	zsetKey := fmt.Sprintf(allTickets, versionMap["poolVersionKey"], matchInfo)
	score, err := redis.Int(redisConn.Do("ZSCORE", zsetKey, playerId))
	if err != nil {
		if err != redis.ErrNil {
			return nil, err
		} else {
			score, err = redis.Int(redisConn.Do("ZSCORE", oldzsetKey, playerId))
			if err != nil {
				if err != redis.ErrNil {
					return nil, err
				} else {
					redisConn.Do("DEL", key)
					return &match_frontend.MatchInfo{}, nil
				}
			}
		}
	}
	return &match_frontend.MatchInfo{
		PlayerId: playerId,
		Score:    int64(score),
		GameId:   gameId,
		SubType:  subType,
	}, nil
}
