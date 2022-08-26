package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nachohotz/go-gorm-restapi/db"
	"github.com/nachohotz/go-gorm-restapi/db/models"
	"github.com/nachohotz/go-gorm-restapi/routes"
)

func main() {
  db.LoadEnvs()
  db.DBConnect()
  db.DB.AutoMigrate(&models.User{})
  db.DB.AutoMigrate(&models.Task{})

  r := mux.NewRouter()

  r.HandleFunc("/", routes.HomeHandler)

  // User routes
  r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
  r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
  r.HandleFunc("/users", routes.PostUsersHandler).Methods("POST")
  r.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")

  // Task routes
  r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
  r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
  r.HandleFunc("/tasks", routes.PostTaskHandler).Methods("POST")
  r.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

  http.ListenAndServe(":3001", r)
}
