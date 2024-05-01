package entity

import "time"

type Cart struct {
	ID                     uint            `gorm:"primary_key" json:"id"`
	Enabled                bool            `json:"enabled"`
	UniqueIdentifier       string          `gorm:"index" json:"unique_identifier"`
	NetTotal               float64         `gorm:"index" json:"net_total"`
	User                   User            `json:"user"`
	CartProducts           []CartProduct   `json:"cart_products"`
	MinBasketLimitViolated bool            `gorm:"-" json:"min_basket_limit_violated"`
	DiscountUsages         []DiscountUsage `json:"discount_usages"`
	UserID                 uint            `json:"-"`
	ShippingMethodNeeded   bool            `gorm:"-" json:"shipping_method_needed"`
	CreatedAt              time.Time       `gorm:"index" json:"created_at"`
	UpdatedAt              time.Time       `gorm:"index" json:"updated_at"`
	DeletedAt              *time.Time      `gorm:"index" json:"-"`
}

