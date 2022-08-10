package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost port=5432 user=nacho dbname=gorm password=sempron"
var DB *gorm.DB
var err error

func DBConnect() {
  // connect to db
  DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

  if err != nil {
    log.Fatal(err)
  } else {
    log.Println("Connected to DB")
  }
}
