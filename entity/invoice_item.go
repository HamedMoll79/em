package entity

import "time"

type InvoiceItem struct {
	ID                  uint               `gorm:"primary_key" json:"id"`
	ProductVariant      ProductVariant     `json:"product_variant"`
	BookingAttributes   *BookingAttributes `gorm:"type:jsonb" json:"booking_attributes"`
	BookingResult       *BookingResult     `gorm:"type:jsonb" json:"booking_result"`
	Invoice             Invoice            `json:"-"`
	NoOfItems           int                `json:"no_of_items"`
	SingleItemPrice     float64            `gorm:"index" json:"single_item_price"`
	TotalItemsPrice     float64            `gorm:"index" json:"total_items_price"`
	VariantAttributes   interface{}        `gorm:"-" json:"variant_attributes"`
	Name                interface{}        `gorm:"-" json:"name"`
	Image               interface{}        `gorm:"-" json:"image"`
	FormAttributes      string             `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"  json:"-"`
	FormAttributesArray interface{}        `gorm:"-" json:"form_attributes"`
	ProductVariantID    uint               `json:"-"`
	InvoiceID           uint               `json:"-"`
	CreatedAt           time.Time          `gorm:"index" json:"created_at"`
	UpdatedAt           time.Time          `gorm:"index" json:"updated_at"`
	DeletedAt           *time.Time         `gorm:"index" json:"-"`
}
