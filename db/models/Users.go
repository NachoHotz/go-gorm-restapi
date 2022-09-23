// Package models ...
package models

import "gorm.io/gorm"

// User ...
type User struct {
	gorm.Model

	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Email     string `json:"email" gorm:"not null;unique_index"`
	Tasks     []Task `json:"tasks"`
}
