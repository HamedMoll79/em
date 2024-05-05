package event_service

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gitlab.sazito.com/sazito/event_publisher/entity"
)

type Service struct {
	RedisDB   *redis.Client
	EventRepo eventRepository
}

type eventRepository interface {
	Save(ctx context.Context, order entity.Event) (entity.Event, error)
	ModifyIsPublished(ctx context.Context, order entity.Event, isPublished bool) (entity.Event, error)
}

func New(redisDB *redis.Client, eventRepo eventRepository) Service {
	return Service{
		RedisDB:   redisDB,
		EventRepo: eventRepo,
	}
}
