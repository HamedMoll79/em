package entity

import "time"

type CartProduct struct {
	ID                  uint               `gorm:"primary_key" json:"id"`
	Cart                Cart               `json:"cart"`
	ProductVariant      ProductVariant     `json:"product_variant"`
	NoOfItems           int                `json:"no_of_items"`
	SingleItemPrice     float64            `gorm:"index" json:"single_item_price"`
	TotalItemsPrice     float64            `gorm:"index" json:"total_items_price"`
	ProductVariantID    uint               `json:"product_variant_id"`
	CartID              uint               `json:"-"`
	FormAttributes      string             `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"  json:"-"`
	FormAttributesArray interface{}        `gorm:"-" json:"form_attributes"`
	BookingAttributes   *BookingAttributes `gorm:"type:jsonb" json:"booking_attributes"`
	ProductName         interface{}        `gorm:"-" json:"product_name"`
	ProductImage        interface{}        `gorm:"-" json:"product_image"`
	CreatedAt           time.Time          `gorm:"index" json:"created_at"`
	UpdatedAt           time.Time          `gorm:"index" json:"updated_at"`
}
