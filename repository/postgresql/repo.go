package postgresql

import "gitlab.sazito.com/sazito/event_publisher/pkg/postgresql"

type PgRepo struct {
	withTestID bool
}

func NewPgRepo(withTestID bool) *PgRepo {
	return &PgRepo{withTestID: withTestID}
}

func (p PgRepo) NewOrdersRepository(getter postgresql.DataContextGetter) *ordersRepository {
	return NewOrdersRepository(getter)
}
