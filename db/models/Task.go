package models

import "gorm.io/gorm"

type Task struct {
  gorm.Model

  Title string `json:"title" gorm:"not null"`
  Description string `json:"description"`
  Done bool `json:"completed"`
  UserID string `json:"userID"`
}
