package redis_adapter

import (
	"github.com/redis/go-redis/v9"
)

type Adapter struct {
	client *redis.Client
}

func New(config Config) Adapter {
	c := redis.NewClient(&redis.Options{
		Username: config.UserName,
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       config.DB,
	})

	return Adapter{client: c}
}

func (a Adapter) Client() *redis.Client {
	return a.client
}
