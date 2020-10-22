package database

import (
	"time"
)

type (
	DB struct {
		URL          string
		Username     string
		DataBaseName string
		MinPoolSize  uint64
		MaxPoolSize  uint64
		Timeout      time.Duration
		PageSize     int64
	}
	Redis struct {
		RedisURL      string
		RedisPassword string
		Timeout       time.Duration
		PoolSize      int
		DialTimeout   time.Duration
	}
)
