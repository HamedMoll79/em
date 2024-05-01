package entity

import (
	"time"
)

type QueueEvent struct {
	ID     uint           `json:"id"`
	Meta   QueueEventMeta `json:"meta"`
	Entity interface{}    `json:"entity"`
}

type QueueEventMeta struct {
	ID         uint            `json:"id"`
	Type       QueueEventType  `json:"type"`
	Date       time.Time       `json:"date"`
	AgentID    uint            `json:"agent_id"`
	ActionType EventActionType `json:"action_type"`
}

type QueueEventType int8

const (
	OrderEntityType QueueEventType = iota + 1
)

type EventActionType int8

const (
	EventCreateType EventActionType = iota + 1
	EventUpdateType
	EventDeleteType
	EventGetType
)
