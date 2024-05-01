package entity

import "time"

type OrderComment struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Order     Order     `json:"order"`
	Comment   string    `json:"comment"`
	OrderID   uint      `json:"-"`
	CreatedAt time.Time `gorm:"index" json:"created_at"`
	UpdatedAt time.Time `gorm:"index" json:"updated_at"`
}
