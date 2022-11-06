package db

import (
  "log"
  "os"

  "github.com/subosito/gotenv"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

var dbPort string
var dbHost string
var dbUser string
var dbPass string
var dbName string

var DSN string

var DB *gorm.DB
var err error

func LoadEnvs() {
  err := gotenv.Load()

  if err != nil {
    log.Fatal("Error loading .env file")
  }

  dbPort = os.Getenv("DB_PORT")
  dbHost = os.Getenv("DB_HOST")
  dbUser = os.Getenv("DB_USER")
  dbPass = os.Getenv("DB_PASSWORD")
  dbName = os.Getenv("DB_NAME")

  DSN = "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " dbname=" + dbName + " password=" + dbPass

  log.Println("Loaded .env file")
}

func Connect() {
  DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

  if err != nil {
    log.Fatal(err)
  } else {
    log.Println("Connected to DB")
  }
}
