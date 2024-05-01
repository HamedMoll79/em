package entity

import "time"

type DiscountCode struct {
	ID                    uint              `gorm:"primary_key" json:"id"`
	Code                  string            `gorm:"index" json:"code"`
	DiscountCondition     DiscountCondition `json:"discount_condition,omitempty"`
	DiscountConditionID   uint              `json:"-"`
	DiscountUsages        []DiscountUsage   `json:"discount_usages"`
	UsageCount            int               `gorm:"-" json:"usage_count"`
	TotalDiscount         float64           `gorm:"-" json:"total_discount"`
	TotalPurchase         float64           `gorm:"-" json:"total_purchase"`
	ShareableLink         string            `gorm:"-" json:"shareable_link"`
	ParentID              uint              `gorm:"default:0" json:"parent_id"`
	DiscountType          int               `gorm:"default:0" json:"type"`
	DiscountSettings      string            `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"  json:"-"`
	DiscountSettingsArray interface{}       `json:"discount_settings" gorm:"-"`
	CreatedAt             time.Time         `gorm:"index" json:"created_at"`
	UpdatedAt             time.Time         `gorm:"index" json:"updated_at"`
	DeletedAt             *time.Time        `gorm:"index" json:"-"`
}
