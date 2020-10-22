package cache

import (
	"context"
	"github.com/poc-redis-lambda-golang-terraform-aws/aws/database"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// CacheHandler representation basic actions for cache
type Handler interface {
	Set(context.Context, string, string) error
	Get(context.Context, string) (string, error)
}

// NewCache creates and returns a new cache for redis instance
func NewCacheRedis(
	redisDB *redis.Client,
	redisDBConfig database.Redis,
	logger *zerolog.Logger) Handler {
	logger.Debug().Msg("New instance cache Redis storage")
	cache := &cacheRedis{
		redisDBConnection: redisDB,
		redisDBConfig:     &redisDBConfig,
		logger:            logger,
	}

	return cache
}

type cacheRedis struct {
	redisDBConfig     *database.Redis
	redisDBConnection *redis.Client
	logger            *zerolog.Logger
}

func (s *cacheRedis) Set(ctx context.Context, key, value string) error {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), s.redisDBConfig.Timeout)
		defer cancel()
	}
	err := s.redisDBConnection.Set(ctx, key, value, 0).Err()

	if err != nil {
		return errors.Wrapf(err, "Error when it set document on redis. Key: [%s] ", key)
	}
	return nil
}

func (s *cacheRedis) Get(ctx context.Context, key string) (string, error) {
	s.logger.Debug().Msgf("Getting object from redis [%s]", key)
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), s.redisDBConfig.Timeout)
		defer cancel()
	}
	val, err := s.redisDBConnection.Get(ctx, key).Result()
	if err == redis.Nil {
		s.logger.Debug().Msgf("key does not exist [%s]", key)
		return "", err
	} else if err != nil {
		return "", err
	}

	return val, nil
}
