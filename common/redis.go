package common

import (
	"context"
	"github.com/go-redis/redis/v7"
	"time"

	//"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

type RedisStorageProvider interface {
	GetKey(ctx context.Context, key string) (string, error)
	SetKey(key string, value interface{}, expireTime time.Duration) error
}

type redisStorageProvider struct {
	client  *redis.Client
	context context.Context
	logger  *logrus.Logger
}

func NewRedisStorageProvider(ctx context.Context, client *redis.Client, logger *logrus.Logger) *redisStorageProvider {
	return &redisStorageProvider{
		client:  client,
		context: ctx,
		logger:  logger,
	}
}

func (redis *redisStorageProvider) GetKey(ctx context.Context, key string) (string, error) {
	val, err := redis.client.Get(key).Result()
	if err != nil {
		redis.logger.Errorf("Get Redis error:%s", err.Error())
	}
	return val, err
}

func (redis *redisStorageProvider) SetKey(key string, value interface{}, expireTime time.Duration) error {
	err := redis.client.Set(key, value, expireTime).Err()
	if err != nil {
		redis.logger.Errorf("Set Redis error:%s", err.Error())
	}
	return err
}
