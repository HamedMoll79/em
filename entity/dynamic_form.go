package entity

import "time"

type DynamicForm struct {
	ID          uint        `gorm:"primary_key" json:"id"`
	Identifier  string      `json:"identifier"`
	Fields      string      `sql:"type:JSONB NOT NULL DEFAULT '[]'::JSONB" json:"-"`
	FieldsMap   interface{} `gorm:"-" json:"fields"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Type        int         `json:"type"`
	Status      string      `json:"status"`
	CreatedAt   time.Time   `gorm:"index" json:"created_at"`
	UpdatedAt   time.Time   `gorm:"index" json:"updated_at"`
	DeletedAt   *time.Time  `gorm:"index" json:"-"`
}
