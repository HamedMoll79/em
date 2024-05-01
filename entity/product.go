package entity

import "time"

type Product struct {
	ID                     uint              `gorm:"primary_key" json:"id"`
	ThemeConfig            string            `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
	ThemeConfigMap         interface{}       `gorm:"-" json:"theme_config"`
	Enabled                bool              `json:"enabled"`
	Name                   string            `json:"name"`
	NameIndex              string            `json:"-"`
	ProductType            int               `json:"-"`
	ProductTypeString      string            `gorm:"-" json:"product_type"`
	StaticUrl              bool              `gorm:"index" json:"static_url"`
	UrlKey                 string            `gorm:"index" json:"-"`
	Url                    string            `gorm:"-" json:"url"`
	ShortUrl               string            `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
	ProductAttributes      string            `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"  json:"-"`
	ProductAttributesArray interface{}       `gorm:"-" json:"product_attributes"`
	Summary                interface{}       `gorm:"-" json:"summary"`
	ProductVariants        []ProductVariant  `json:"product_variants"`
	ProductCategories      []ProductCategory `gorm:"many2many:products_product_categories;" json:"product_categories"`
	Tags                   []Tag             `gorm:"many2many:products_tags;" json:"tags"`
	Images                 []Image           `gorm:"many2many:products_images;" json:"images"`
	DynamicFormID          *uint             `json:"dynamic_form_id"`
	DynamicForm            *DynamicForm      `json:"dynamic_form,omitempty"`
	EventEntityID          *string           `gorm:"type:varchar(36)" json:"event_entity_id"`
	CoverImage             *Image            `json:"-"`
	CoverImageID           uint              `json:"-"`
	CreatedAt              time.Time         `gorm:"index" json:"created_at"`
	UpdatedAt              time.Time         `gorm:"index" json:"updated_at"`
	DeletedAt              *time.Time        `gorm:"index" json:"-"`
}