package models

import (
	"gorm.io/gorm"
)

type Deliver struct {
	gorm.Model
	ProductId uint64
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Amount    float64
}

/* Set table name. */
func (d *Deliver) TableName() string {
	return "delivers"
}
