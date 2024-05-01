package entity

import "time"

type Tag struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Title     string    `json:"title"`
	Lang      int       `json:"-"`
	Products  []Product `gorm:"many2many:products_tags;" json:"products"`
	CreatedAt time.Time `gorm:"index" json:"created_at"`
	UpdatedAt time.Time `gorm:"index" json:"updated_at"`
}
