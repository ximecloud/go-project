package rds

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string
	Port     string
	Password string
	Database int
}

func New(config *Config) (*redis.Client, error) {
	rds := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password:     config.Password,
		DB:           config.Database,
		PoolSize:     50,
		MinIdleConns: 10,
	})

	ctx := context.Background()
	_, err := rds.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %v", err)
	}

	return rds, nil
}
