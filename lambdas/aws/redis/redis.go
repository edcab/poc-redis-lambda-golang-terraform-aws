package redis

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisDataBase struct {
	databaseConnection *redis.Client
}

// NewRedisDBStorage creates and returns redis db connection instance
func NewRedisDBStorage(ctx context.Context) (*redisDataBase, error) {
	log.Println("NewRedisDBStorage")
	dataBase := &redisDataBase{}
	err := dataBase.OpenConnection(ctx)
	if err != nil {
		log.Println("error opening connection", err)
		return nil, err
	}

	return dataBase, nil
}

// OpenConnection start redis db connection
func (db *redisDataBase) OpenConnection(ctx context.Context) error {
	log.Println("opening connection")
	rdb := redis.NewClient(&redis.Options{
		Addr:        "poc-redis-cluster.sfofiy.0001.use1.cache.amazonaws.com:6379", //hard code, this is better retrieve from parameter store
		Password:    "",
		DB:          0,
		DialTimeout: 30 * time.Second,
	})
	log.Println("rdb: ", rdb)

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("Ping Redis connection: %s", pong)

	db.databaseConnection = rdb
	log.Println("RedisDB UP")
	return nil
}

// GetConnection get redisDB connection
func (db *redisDataBase) GetConnection() *redis.Client {
	return db.databaseConnection
}

// CloseConnection close redisDB connection
func (db *redisDataBase) CloseConnection() {
	log.Printf("Close RedisDB connection (%s)", "localhost:6379")
	err := db.databaseConnection.Close()
	if err != nil {
		log.Println("Error when it did the disconnection to redisDB")
		panic(err)
	}
}
