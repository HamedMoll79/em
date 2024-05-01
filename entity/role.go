package entity

import (
	"github.com/lib/pq"
	"time"
)

type Role struct {
	ID            uint           `gorm:"primary_key" json:"id"`
	Name          string         `gorm:"index" json:"name"`
	Color         string         `gorm:"index" json:"color"`
	Users         []User         `gorm:"many2many:users_roles" json:"users" `
	RoleConfig    pq.StringArray `gorm:"type:varchar(100)[];not null;DEFAULT:array['ALL']::varchar(100)[]" json:"role_config"`
	CreatedAt     time.Time      `gorm:"index" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"index" json:"updated_at"`
	RoleConfigMap map[int]bool   `gorm:"-" json:"-"`
}