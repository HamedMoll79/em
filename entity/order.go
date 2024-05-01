package entity

import "time"

type Order struct {
	ID                     uint           `gorm:"primary_key" json:"id"`
	OrderIdentifier        string         `gorm:"index" json:"order_identifier"`
	OrderStatus            int            `gorm:"index" json:"-"`
	PaymentStatus          *int           `gorm:"index" json:"-"`
	Seen                   bool           `gorm:"index" json:"seen"`
	User                   User           `json:"user"`
	Receipt                Receipt        `json:"receipt"`
	Invoice                Invoice        `json:"invoice"`
	OrderComments          []OrderComment `json:"order_comments"`
	InvoiceID              uint           `json:"-"`
	ReceiptID              uint           `json:"-"`
	UserID                 uint           `json:"-"`
	Pinched                bool           `gorm:"index;default:false" json:"-"`
	PinchedAt              *time.Time     `gorm:"index" json:"-"`
	ShippingTrackingDetail string         `json:"shipping_tracking_detail"`
	CompleteCustomMessage  string         `json:"complete_custom_message"`
	ShippingDetails        string         `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
	ShippingDetailsObject  interface{}    `gorm:"-" json:"shipping_details"`
	OrderNumber            interface{}    `gorm:"-" json:"order_number"`
	Note                   string         `gorm:"note" json:"note"`
	AddNoteToPdf           bool           `gorm:"add_note_to_pdf" json:"add_note_to_pdf"`
	Status                 interface{}    `gorm:"-" json:"status"`
	StatusTitle            interface{}    `gorm:"-" json:"status_title"`
	CreatedAt              time.Time      `gorm:"index" json:"created_at"`
	UpdatedAt              time.Time      `gorm:"index" json:"updated_at"`
	DeletedAt              *time.Time     `gorm:"index" json:"-"`
}
