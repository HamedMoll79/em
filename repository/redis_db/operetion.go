package redis_db

import (
	"context"
	"gitlab.sazito.com/sazito/event_publisher/pkg/richerror"
)

func (d *DB)FetchMessage(ctx context.Context, queueName string) (string, error) {
	const op = "redis_adapter.FetchMessage"
	data, err := d.conn.LPop(ctx, queueName).Result()
	if err != nil {

		return "", richerror.New(op).WithOp(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return data, nil
}
