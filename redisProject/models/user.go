package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Email    string `gorm:"column:email;unique"`
	Password string `gorm:"column:password"`
}
