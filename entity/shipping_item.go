package entity

import "time"

type ShippingItem struct {
	ID                     uint                              `gorm:"primary_key" json:"id"`
	InvoiceID              uint                              `json:"invoice_id"`
	InvoiceItemIDs         string                            `json:"invoice_item_ids"`
	ShippingStatus         int                               `gorm:"index" json:"-"`
	Rate                   string                            `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
	RateArray              interface{}                       `gorm:"-" json:"rate"`
	ShippingNumber         interface{}                       `gorm:"-" json:"shipping_number"`
	ShippingTrackingDetail string                            `json:"shipping_tracking_detail"`
	CompleteCustomMessage  string                            `json:"complete_custom_message"`
	Status                 interface{}                       `gorm:"-" json:"status"`
	StatusTitle            interface{}                       `gorm:"-" json:"status_title"`
	CreatedAt              time.Time                         `gorm:"index" json:"created_at"`
	UpdatedAt              time.Time                         `gorm:"index" json:"updated_at"`
	ThirdPartyDetails      ThirdPartyShipmentTrackingDetails `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"third_party_details"`
}

type ThirdPartyShipmentTrackingDetails struct {
	Items []ThirdPartyShipmentTrackingDetail `json:"items"`
}

type ThirdPartyShipmentTrackingDetail struct {
	ThirdParty  string      `json:"thirdParty"`
	SubmittedAt time.Time   `json:"submittedAt"`
	Payload     interface{} `json:"payload"`
}
