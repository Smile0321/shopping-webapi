package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `gorm:"unique; size:255; NOT NULL"`
	Password     string `gorm:"size:255; NOT NULL"`
	FirstName    string `gorm:"size:255; NOT NULL"`
	LastName     string `gorm:"size:255; NOT NULL"`
	Role         int8
	Permission   bool
	RefreshToken string
}

/* Set table name. */
func (u *User) TableName() string {
	return "users"
}

/* Hash password. */
func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil
	}
	u.Password = string(bytes)
	return nil
}

/* Verify hashed password. */
func (u *User) VerifyPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
