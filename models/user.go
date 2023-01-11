package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string `json:"full_name" gorm:"not null; index: idx_name;"`
	Email    string `json:"email" gorm:"unique;not null;"`
	Password string `json:"password" gorm:"type:text;not null"`
	Avatar   string `json:"avatar" gorm:"default:''"`
}
