package redisdb

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

type redisDataBase struct {
	databaseConnection *redis.Client
}

// NewRedisDBStorage creates and returns redis db connection instance
func NewRedisDBStorage(ctx context.Context) (*redisDataBase, error) {

	dataBase := &redisDataBase{}
	err := dataBase.OpenConnection(ctx)
	if err != nil {
		return nil, err
	}

	return dataBase, nil
}

// OpenConnection start redis db connection
func (db *redisDataBase) OpenConnection(ctx context.Context) error {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.Printf("Ping Redis connection: %s", pong)
	fmt.Printf("Ping Redis connection: %s", pong)

	db.databaseConnection = rdb
	log.Println("RedisDB UP")
	fmt.Println("RedisDB UP")
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
