package entity

import "time"

type Receipt struct {
	ID            uint       `gorm:"primary_key" json:"id"`
	ReceiptRef    string     `json:"receipt_ref"`
	ImageUrl      string     `json:"image_url"`
	Invoice       Invoice    `json:"-"`
	Payment       Payment    `json:"payment"`
	ReceiptAmount float64    `gorm:"index" json:"receipt_amount"`
	InvoiceID     uint       `json:"-"`
	PaymentID     uint       `json:"-"`
	CreatedAt     time.Time  `gorm:"index" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"index" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"index" json:"-"`
}
