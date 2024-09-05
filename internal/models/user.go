package models

import (
	"gorm.io/gorm"
)

// User struct
//
// swagger:model
// Defines the structure of a User
type User struct {
	gorm.Model
	Email     string `json:"email" gorm:"unique" binding:"required"`
	Password  string `json:"password" binding:"required"`
	UserID    string `json:"user_id" gorm:"unique"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}
