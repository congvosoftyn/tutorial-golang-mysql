package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `json:"title" gorm:"type:varchar(191);not null"`
	Description string `json:"description" gorm:"type:text;not null"`
	Author      string `json:"author"`
	Price       uint   `json:"price" gorm:"not null"`
}
