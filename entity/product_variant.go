package entity

import "time"

type ProductVariant struct {
	ID                     uint             `gorm:"primary_key" json:"id"`
	Enabled                bool             `json:"enabled"`
	Sku                    string           `gorm:"index" json:"sku"`
	StockNumber            int              `json:"stock_number"`
	IsStockManaged         bool             `gorm:"index" json:"is_stock_managed"`
	Price                  float64          `gorm:"index" json:"price"`
	RawPrice               float64          `gorm:"index" json:"-"`
	RawPriceWithNull       interface{}      `gorm:"-" json:"raw_price"`
	HasRawPrice            bool             `json:"-"`
	MaxNoOfOrder           int              `json:"max_no_of_order"`
	HasMaxNoOfOrder        bool             `json:"has_max_order"`
	MinOrderCount          int              `json:"min_order_count"`
	Weight                 float64          `gorm:"index" json:"weight"`
	ProductAttributes      string           `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
	SortIndex              int              `sql:"type:int; DEFAULT:0" json:"sort_index"`
	ProductAttributesArray interface{}      `gorm:"-" json:"product_attributes"`
	Product                Product          `json:"product"`
	ProductID              uint             `json:"-"`
	Image                  *Image           `json:"-"`
	ImageID                *uint            `json:"image_id"`
	CommercialFiles        []CommercialFile `gorm:"many2many:product_variants_commercial_files;" json:"commercial_files"`
	Status                 string           `json:"status" gorm:"-"`
	CreatedAt              time.Time        `gorm:"index" json:"created_at"`
	UpdatedAt              time.Time        `gorm:"index" json:"updated_at"`
	DeletedAt              *time.Time       `gorm:"index" json:"-"`
	SoldCount              int              `gorm:"index" json:"sold_count"`
}
