package postgresql

import (
	"context"
	"database/sql"
	"gitlab.sazito.com/sazito/event_publisher/entity"
	"gitlab.sazito.com/sazito/event_publisher/pkg/postgresql"
	"gitlab.sazito.com/sazito/event_publisher/pkg/richerror"
	"gitlab.sazito.com/sazito/event_publisher/pkg/utils"
)

type eventsRepository struct {
	conn *sql.DB
}

func NewEventsRepository(getter postgresql.DataContextGetter) *eventsRepository {
	return &eventsRepository{
		conn: getter.GetDataContext(),
	}
}

func (d *eventsRepository) Save(ctx context.Context, event entity.Event) (entity.Event, error) {
	const op = "postgresql.eventsRepository.Save"

	id := ""
	err := d.conn.QueryRowContext(ctx, `INSERT INTO events(entity_type, action_type, agent_id, date, entity, store_id) VALUES ($1 , $2, $3, $4, $5) RETURNING id`,
		event.EntityType, event.ActionType, event.AgentID, event.Date, event.Entity, event.StoreID).Scan(&id)
	if err != nil {
		return entity.Event{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}


	event.ID = utils.ParseUint(id)

	return event, nil
}

func (d *eventsRepository) ModifyIsPublished(ctx context.Context, order entity.Event, isPublished bool) (entity.Event, error) {
	const op = "postgresql.eventsRepository.ModifyIsPublished"

	err := d.conn.QueryRowContext(ctx, `UPDATE events SET is_published = $1 WHERE id = $2 RETURNING id`, isPublished, order.ID)
	if err != nil {

		return entity.Event{}, richerror.New(op).WithErr(err.Err())
	}

	order.IsPublished = isPublished

	return order, nil
}
