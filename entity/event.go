package entity

import (
	"time"
)

type Event struct {
	ID          uint        `json:"id"`
	EntityType  EntityType  `json:"entity_type"`
	ActionType  ActionType  `json:"action_type"`
	AgentID     uint        `json:"agent_id"`
	Date        time.Time   `json:"date"`
	IsPublished bool        `json:"is_published"`
	StoreID     uint        `json:"store_id"`
	Entity      interface{} `json:"entity"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type EntityType int8

const (
	OrderEntityType EntityType = iota + 1
)

type ActionType int8

const (
	EventCreateType ActionType = iota + 1
	EventUpdateType
	EventDeleteType
	EventGetType
)
