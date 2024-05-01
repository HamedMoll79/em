package entity

import "time"

type ProductCategory struct {
	ID              uint        `gorm:"primary_key" json:"id"`
	ThemeConfig     string      `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
	ThemeConfigMap  interface{} `gorm:"-" json:"theme_config"`
	Name            string      `json:"name"`
	NameIndex       string      `json:"-"`
	Enabled         bool        `json:"enabled"`
	StaticUrl       bool        `gorm:"index" json:"static_url"`
	Url             string      `gorm:"-" json:"url"`
	ShortUrl        string      `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
	UrlKey          string      `gorm:"index" json:"-"`
	Attributes      string      `sql:"type:JSONB NOT NULL DEFAULT '[]'::JSONB"  json:"-"`
	AttributesArray interface{} `gorm:"-" json:"attributes"`
	Description     string      `json:"description"`
	ProductsCount   interface{} `gorm:"-" json:"products_count"`
	Products        []Product   `json:"products"`
	CreatedAt       time.Time   `gorm:"index" json:"created_at"`
	UpdatedAt       time.Time   `gorm:"index" json:"updated_at"`
	DeletedAt       *time.Time  `gorm:"index" json:"-"`
}
