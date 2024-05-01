package entity

import "time"

type Payment struct {
	ID                     uint          `gorm:"primary_key" json:"id"`
	PaymentStatus          int           `json:"-"`
	PaymentStatusInterface interface{}   `gorm:"-" json:"payment_status"`
	PaymentStatusFa        string        `gorm:"-" json:"payment_status_fa"`
	PaymentAmount          float64       `gorm:"index" json:"payment_amount"`
	IsActive               bool          `json:"is_active"`
	PaymentIdentifier      string        `gorm:"index" json:"payment_identifier"`
	Invoice                Invoice       `json:"-"`
	PaymentType            int           `json:"-"`
	PaymentTypeMap         PaymentType   `gorm:"-" json:"payment_type"`
	PaymentSubType         int           `json:"-"`
	PaymentSubTypeString   string        `gorm:"-" json:"payment_sub_type"`
	PaymentSteps           []PaymentStep `json:"-"`
	InvoiceID              uint          `json:"invoice_id"`
	CreatedAt              time.Time     `gorm:"index" json:"created_at"`
	UpdatedAt              time.Time     `gorm:"index" json:"updated_at"`
	DeletedAt              *time.Time    `gorm:"index" json:"-"`
}

type PaymentType struct {
	Id            int                    `json:"id"`
	Title         string                 `json:"title"`
	TitleFa       string                 `json:"title_fa"`
	Description   map[string]interface{} `json:"description"`
	ReferenceCode string                 `json:"reference_code"`
	Order         int                    `json:"order"`
	IsDefault     bool                   `json:"is_default"`
}

type PaymentStep struct {
	ID              uint       `gorm:"primary_key" json:"id"`
	Payment         Payment    `json:"payment"`
	StepPaymentInfo string     `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"step_payment_info"`
	StepStatus      int        `json:"-"`
	StepNumber      int        `json:"step_number"`
	StepType        int        `json:"step_type"`
	PaymentID       uint       `json:"-"`
	CreatedAt       time.Time  `gorm:"index" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"index" json:"updated_at"`
	DeletedAt       *time.Time `gorm:"index" json:"-"`
}