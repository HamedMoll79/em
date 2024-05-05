package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
	"gitlab.sazito.com/sazito/event_publisher/entity"
	"gitlab.sazito.com/sazito/event_publisher/pkg/postgresql"
	"gitlab.sazito.com/sazito/event_publisher/pkg/richerror"
)

type webhooksRepository struct {
	conn *sql.DB
}

func NewWebhooksRepository(getter postgresql.DataContextGetter) *webhooksRepository {
	return &webhooksRepository{
		conn: getter.GetDataContext(),
	}
}

func (d *webhooksRepository) GetByStoreID(ctx context.Context, storeID uint) (entity.Webhook, error) {
	const op = "postgresql.GetByStoreID"

	var message []byte

	err := d.conn.QueryRowContext(ctx, `SELECT (id, url, store_id) FROM webhooks WHERE store_id = $1`, storeID).Scan(&message)
	if err != nil {
		if err == sql.ErrNoRows {

			return entity.Webhook{}, richerror.New(op).WithErr(err).WithKind(richerror.KindNotFound)
		}

		return entity.Webhook{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	var webhook entity.Webhook
	err = json.Unmarshal(message, &webhook)
	if err != nil {
		return entity.Webhook{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return webhook, nil
}
