package entity

import "time"

type Image struct {
	ID              uint      `gorm:"primary_key" json:"id"`
	Name            string    `json:"name"`
	Alt             string    `json:"alt"`
	Width           int       `json:"width"`
	Height          int       `json:"height"`
	WidthRatio      float64   `json:"width_ratio"`
	HeightRatio     float64   `json:"height_ratio"`
	RootImage       string    `json:"-"`
	Url             string    `gorm:"-" json:"url"`
	Thumb           string    `gorm:"-" json:"thumb"`
	Path            string    `gorm:"-" json:"-"`
	ImageOrder      int       `json:"order"`
	ImageIdentifier string    `json:"-"`
	IsProtected     bool      `gorm:"default:false" json:"-"`
	CreatedAt       time.Time `gorm:"index" json:"created_at"`
	UpdatedAt       time.Time `gorm:"index" json:"updated_at"`
	IsProductImage  bool      `gorm:"-" json:"-"`
}

