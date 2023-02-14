package models

import (
	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model
	ProductId uint64
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Amount    float64
}

/* Set table name. */
func (p *Purchase) TableName() string {
	return "purchases"
}
