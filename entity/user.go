package entity

import "time"

type User struct {
	ID                           uint              `gorm:"primary_key" json:"id"`
	Email                        string            `gorm:"index" json:"email"`
	MobilePhone                  string            `gorm:"index" json:"mobile_phone"`
	MobilePhoneVerifiedAt        *time.Time        `json:"mobile_phone_verified_at"`
	FirstName                    string            `gorm:"index" json:"first_name"`
	LastName                     string            `gorm:"index" json:"last_name"`
	Password                     string            `gorm:"-" json:"password,omitempty" `
	PasswordDigest               string            `json:"-"`
	ForgotPasswordToken          string            `json:"-"`
	ForgotPasswordTokenCreatedAt *time.Time        `json:"-"`
	ShippingAddresses            []ShippingAddress `json:"shipping_addresses"`
	Orders                       []Order           `json:"-"`
	Carts                        []Cart            `json:"-"`
	Roles                        []Role            `gorm:"many2many:users_roles;" json:"roles"`
	Sold                         interface{}       `gorm:"-" json:"sold"`
	IsGuest                      bool              `gorm:"-" json:"is_guest"`
	IsGod                        bool              `gorm:"-" json:"is_god"`
	SoldAmount                   interface{}       `gorm:"-" json:"sold_amount"`
	TokenRateLimit               int               `json:"-"`
	TokenRateLimitedAt           time.Time         `json:"-"`
	CreatedAt                    time.Time         `gorm:"index" json:"created_at"`
	UpdatedAt                    time.Time         `gorm:"index" json:"updated_at"`
	DeletedAt                    *time.Time        `gorm:"index" json:"-"`
}
