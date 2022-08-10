package models

import "gorm.io/gorm"

type User struct {
  gorm.Model

  ID int `json:"id"`
  FirstName string `json:"first_name" gorm:"not null"`
  LastName string `json:"last_name" gorm:"not null"`
  Email string `json:"email" gorm:"not null;unique_index"`
  Tasks []Task `json:"tasks"`
}
