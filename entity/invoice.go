package entity

import "time"

type Invoice struct {
	ID                          uint            `gorm:"primary_key" json:"id"`
	InvoiceIdentifier           string          `gorm:"index" json:"invoice_identifier"`
	NetTotal                    float64         `gorm:"index" json:"net_total"`
	FinalTotal                  float64         `gorm:"index" json:"final_total"`
	DiscountTotal               float64         `gorm:"index" json:"discount_total"`
	ShippingTotal               float64         `gorm:"index" json:"shipping_total"`
	Vat                         float64         `gorm:"index" json:"vat"`
	VatPercent                  float64         `json:"vat_percent"`
	User                        User            `json:"user"`
	ShippingAddress             ShippingAddress `json:"shipping_address"`
	ShippingMethod              int             `json:"-"`
	ShippingMethodString        interface{}     `gorm:"-" json:"shipping_method"`
	ShippingItems               []ShippingItem  `json:"-"`
	ShippingItemArray           interface{}     `gorm:"-" json:"shipping_items"`
	DiscountUsages              []DiscountUsage `json:"discount_usages"`
	InvoiceItems                []InvoiceItem   `json:"invoice_items"`
	Payments                    []Payment       `json:"-"`
	Receipts                    []Receipt       `json:"receipts"`
	Orders                      []Order         `json:"-"`
	UserID                      uint            `json:"-"`
	ShippingAddressID           uint            `json:"-"`
	ShippingMethodNeeded        bool            `gorm:"-" json:"shipping_method_needed"`
	ShippingDetailsCustomer     string          `gorm:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
	ShippingDetailsCustomerData interface{}     `gorm:"-" json:"shipping_details_customer"`
	UserComment                 string          `json:"user_comment"`
	CreatedAt                   time.Time       `gorm:"index" json:"created_at"`
	UpdatedAt                   time.Time       `gorm:"index" json:"updated_at"`
	DeletedAt                   *time.Time      `gorm:"index" json:"-"`
}