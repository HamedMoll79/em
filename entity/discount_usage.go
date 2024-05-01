package entity

import "time"

type DiscountUsage struct {
	ID                   uint         `gorm:"primary_key" json:"id"`
	DiscountCode         DiscountCode `json:"discount_code"`
	DiscountCodeID       uint         `json:"-"`
	DiscountCodeParentID uint         `json:"-"`
	Invoice              Invoice      `json:"invoice"`
	InvoiceId            uint         `json:"-"`
	Used                 bool         `json:"used"`
	CreatedAt            time.Time    `gorm:"index" json:"created_at"`
	UpdatedAt            time.Time    `gorm:"index" json:"updated_at"`
	DeletedAt            *time.Time   `gorm:"index" json:"-"`
}
