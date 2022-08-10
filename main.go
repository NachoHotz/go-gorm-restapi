package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/nachohotz/go-gorm-restapi/db"
	"github.com/nachohotz/go-gorm-restapi/db/models"
	"github.com/nachohotz/go-gorm-restapi/routes"
)

func main() {
  db.DBConnect()
  db.DB.AutoMigrate(models.Task{})
  db.DB.AutoMigrate(models.User{})

  r := mux.NewRouter()

  r.HandleFunc("/", routes.HomeHandler)

  http.ListenAndServe(":3001", r)
}
