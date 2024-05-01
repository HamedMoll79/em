package entity

import (
	"time"
)

type QueueEvent struct {
	Meta   QueueEventMeta `json:"meta"`
	Entity interface{}    `json:"entity"`
}

type QueueEventMeta struct {
	Type       QueueEventType  `json:"type"`
	Date       time.Time       `json:"date"`
	AgentID    uint            `json:"agent_id"`
	ActionType EventActionType `json:"action_type"`
}

type QueueEventType int8

const (
	_               = iota
	OrderEntityType = 1
)

type EventActionType int8

const (
	_               = iota
	EventCreateType = 1
	EventUpdateType = 2
	EventDeleteType = 3
	EventGetType    = 4
)
