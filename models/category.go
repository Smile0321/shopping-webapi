package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `gorm:"size:255; NOT NULL"`
}

/* Set table name. */
func (c *Category) TableName() string {
	return "categories"
}
