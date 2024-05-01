package entity

import "time"

type ShippingAddress struct {
	ID                        uint       `gorm:"primary_key" json:"id"`
	Email                     string     `json:"email"`
	FirstName                 string     `json:"first_name"`
	LastName                  string     `json:"last_name"`
	Address                   string     `json:"address"`
	PostalCode                string     `json:"postal_code"`
	PhoneNumber               string     `json:"phone_number"`
	MobilePhone               string     `json:"mobile_phone"`
	ShippingAddressIdentifier string     `gorm:"index" json:"shipping_address_identifier"`
	Longitude                 float64    `json:"longitude"`
	Latitude                  float64    `json:"latitude"`
	Region                    Region     `json:"region"`
	City                      City       `json:"city"`
	User                      User       `json:"user"`
	UserID                    uint       `json:"user_id"`
	CityID                    uint       `json:"-"`
	RegionID                  uint       `json:"-"`
	ShowMap                   bool       `json:"show_map"`
	Invoices                  []Invoice  `json:"-"`
	CreatedAt                 time.Time  `gorm:"index" json:"created_at"`
	UpdatedAt                 time.Time  `gorm:"index" json:"updated_at"`
	DeletedAt                 *time.Time `gorm:"index" json:"-"`
}

type Region struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      string     `gorm:"index" json:"name"`
	Longitude float64    `json:"longitude"`
	Latitude  float64    `json:"latitude"`
	Cities    []City     `json:"cities"`
	CreatedAt time.Time  `gorm:"index" json:"created_at"`
	UpdatedAt time.Time  `gorm:"index" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
}

type City struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      string     `gorm:"index" json:"name"`
	Longitude float64    `json:"longitude"`
	Latitude  float64    `json:"latitude"`
	Region    Region     `json:"region"`
	RegionID  uint       `json:"-"`
	CreatedAt time.Time  `gorm:"index" json:"created_at"`
	UpdatedAt time.Time  `gorm:"index" json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}
