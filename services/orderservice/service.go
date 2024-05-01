package orderservice

import (
	"context"
	"github.com/adjust/redismq"
	"gitlab.sazito.com/sazito/event_publisher/entity"
)

type Service struct {
	RedisQueue *redismq.Queue
	OrderRepo  OrderRepository
}

type OrderRepository interface {
	Save(ctx context.Context, order entity.Order) (entity.Order, error)
	ModifyIsPublished(ctx context.Context, order entity.Order, isPublished bool) (entity.Order, error)
}
