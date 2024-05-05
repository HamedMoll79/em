package redis_db

import (
	"github.com/redis/go-redis/v9"
)

type DB struct {
	conn *redis.Conn
}

func (m *DB) Conn() *redis.Conn {
	return m.conn
}

func New(conn *redis.Conn) *DB {
	return &DB{conn: conn}
}
