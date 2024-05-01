package redisqueue

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gitlab.sazito.com/sazito/event_publisher/pkg/richerror"
)

var c *redis.Client

func New(config Config) *redis.Client {
	c = redis.NewClient(&redis.Options{
		Username: config.UserName,
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       config.DB,
	})

	return c
}

func FetchMessage(queueName string) (string, error) {
	const op = "redisqueue.FetchMessage"
	data, err := c.LPop(context.Background(), queueName).Result()
	if err != nil {

		return "", richerror.New(op).WithOp(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return data, nil
}
