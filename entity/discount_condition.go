package entity

import "time"

type DiscountCondition struct {
	ID                                   uint              `gorm:"primary_key" json:"id"`
	Enabled                              bool              `json:"enabled"`
	Products                             []Product         `gorm:"many2many:discount_conditions_products;" json:"products"`
	ProductCategories                    []ProductCategory `gorm:"many2many:discount_conditions_product_categrories;" json:"product_categories"`
	DiscountType                         int               `json:"-"`
	DiscountTypeInterface                interface{}       `gorm:"-" json:"discount_type"`
	SoftInclude                          bool              `sql:"type:bool; DEFAULT:false" json:"soft_include"`
	CertainProductsMinQuantity           int               `json:"certain_products_min_quantity"`
	CertainProductsMin                   bool              `sql:"type:bool; DEFAULT:false" json:"certain_products_min"`
	CertainCategoriesProductsMin         bool              `sql:"type:bool; DEFAULT:false" json:"certain_categories_products_min"`
	CertainCategoriesProductsMinQuantity int               `json:"certain_categories_products_min_quantity"`
	DiscountAmount                       float64           `json:"discount_amount"`
	DiscountMaximumAmount                float64           `json:"discount_maximum_amount"`
	AppliesTo                            int               `json:"-"`
	AppliesToInterface                   interface{}       `gorm:"-" json:"applies_to"`
	CartMinimumAmount                    float64           `json:"cart_minimum_amount"`
	CartMinimumQuantity                  float64           `json:"cart_minimum_quantity"`
	CartMaximumAmount                    float64           `json:"cart_maximum_amount"`
	CartMaximumQuantity                  float64           `json:"cart_maximum_quantity"`
	CustomerGroup                        int               `json:"-"`
	CustomerGroupInterface               interface{}       `gorm:"-" json:"customer_group"`
	IsMaxUsageLimited                    bool              `json:"is_max_usage_limited"`
	MaxUsageCount                        int               `json:"max_usage_count"`
	IsMaxUserUsageLimited                bool              `json:"is_max_user_usage_limited"`
	MaxUserUsageCount                    int               `json:"max_user_usage_count"`
	StartDate                            *time.Time        `json:"start_date"`
	EndDate                              *time.Time        `json:"end_date"`
	Users                                []User            `gorm:"many2many:discount_conditions_users;" json:"users"`
	IsExpired                            interface{}       `gorm:"-" json:"is_expired"`
	CreatedAt                            time.Time         `gorm:"index" json:"created_at"`
	UpdatedAt                            time.Time         `gorm:"index" json:"updated_at"`
	DeletedAt                            *time.Time        `gorm:"index" json:"-"`
}
