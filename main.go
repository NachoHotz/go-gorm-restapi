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
  db.DB.AutoMigrate(models.Task{}, models.User{})

  r := mux.NewRouter()

  r.HandleFunc("/", routes.HomeHandler)

  r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
  r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
  r.HandleFunc("/users", routes.PostUsersHandler).Methods("POST")
  r.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")

  http.ListenAndServe(":3001", r)
}
