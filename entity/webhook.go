package entity

import "time"

type Webhook struct {
	ID        uint      `json:"id"`
	Url       string    `json:"url"`
	StoreID   int       `json:"store_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
