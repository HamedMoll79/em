package postgresql

import "gitlab.sazito.com/sazito/event_publisher/pkg/postgresql"

type PgRepo struct {
	withTestID bool
}

func NewPgRepo(withTestID bool) *PgRepo {
	return &PgRepo{withTestID: withTestID}
}

func (p PgRepo) NewEventsRepository(getter postgresql.DataContextGetter) *eventsRepository {
	return NewEventsRepository(getter)
}

func (p PgRepo) NewWebhooksRepository(getter postgresql.DataContextGetter) *webhooksRepository {
	return NewWebhooksRepository(getter)
}
