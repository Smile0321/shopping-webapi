package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	CategoryId uint64
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name       string   `gorm:"size:255; NOT NULL"`
	Standard   string   `gorm:"size:255; NOT NULL"`
	Unit       string   `gorm:"size:50; NOT NULL"`
	Price      float32
	Amount     float64
}

/* Set table name. */
func (p *Product) TableName() string {
	return "products"
}
