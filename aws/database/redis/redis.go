package redisdb

import (
	"context"
	"github.com/poc-redis-lambda-golang-terraform-aws/aws/database"

	"github.com/go-redis/redis/v8"

	"github.com/rs/zerolog"

	"github.com/pkg/errors"
)

type redisDataBase struct {
	databaseConnection *redis.Client
	config             *database.Redis
	logger             *zerolog.Logger
}

// NewRedisDBStorage creates and returns redis db connection instance
func NewRedisDBStorage(configDB *database.Redis, logger *zerolog.Logger) (database.DataBase, error) {
	logger.Debug().Msgf("New instance Redis storage [%s]", configDB.RedisURL)

	dataBase := &redisDataBase{
		config: configDB,
		logger: logger,
	}
	err := dataBase.OpenConnection()
	if err != nil {
		return nil, err
	}
	return dataBase, nil
}

// OpenConnection start redis db connection
func (db *redisDataBase) OpenConnection() error {
	db.logger.Info().Msgf("Starting redisDB connection (%s)", db.config.RedisURL)

	ctx, cancelFunc := context.WithTimeout(context.Background(), db.config.Timeout)
	defer cancelFunc()

	rdb := redis.NewClient(&redis.Options{
		Addr:        db.config.RedisURL,
		Password:    db.config.RedisPassword,
		DB:          0,
		DialTimeout: db.config.DialTimeout,
		PoolSize:    db.config.PoolSize,
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		return errors.Wrap(err, "Error on connection to redis")
	}
	db.logger.Info().Msgf("Ping Redis connection: %s", pong)

	db.databaseConnection = rdb
	db.logger.Info().Msg("RedisDB UP")
	return nil
}

// GetConnection get redisDB connection
func (db *redisDataBase) GetConnection() interface{} {
	return db.databaseConnection
}

// CloseConnection close redisDB connection
func (db *redisDataBase) CloseConnection() {
	db.logger.Debug().Msgf("Close RedisDB connection (%s)", db.config.RedisURL)
	err := db.databaseConnection.Close()
	if err != nil {
		db.logger.Error().Err(err).Msg("Error when it did the disconnection to redisDB")
		panic(err)
	}
}
