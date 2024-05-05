package event_service

import (
	"context"
	"gitlab.sazito.com/sazito/event_publisher/entity"
	"gitlab.sazito.com/sazito/event_publisher/repository/redis_db"
)

type Service struct {
	RedisDB        *redis_db.DB
	EventRepo      eventRepository
	RedisQueueName string
	WebhookRepo    webHookRepository
}

type eventRepository interface {
	Save(ctx context.Context, order entity.Event) (entity.Event, error)
	ModifyIsPublished(ctx context.Context, order entity.Event, isPublished bool) (entity.Event, error)
}

type webHookRepository interface {
	GetByStoreID(ctx context.Context, storeID uint) (entity.Webhook, error)
}

func New(redisDB *redis_db.DB, eventRepo eventRepository, webhookRepo webHookRepository) Service {
	return Service{
		RedisDB:     redisDB,
		EventRepo:   eventRepo,
		WebhookRepo: webhookRepo,
	}
}
