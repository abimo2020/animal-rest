package models

import "gorm.io/gorm"

type Animal struct {
	gorm.Model
	Name  string `json:"name" form:"name" gorm:"unique"`
	Class string `json:"class" form:"class"`
	Legs  uint   `json:"legs" form:"legs"`
}
