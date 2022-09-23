// Package main
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nachohotz/go-gorm-restapi/db"
	"github.com/nachohotz/go-gorm-restapi/db/models"
	"github.com/nachohotz/go-gorm-restapi/routes"
)

func mountRouter(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(handler)
}

func main() {
	db.LoadEnvs()
	db.Connect()
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.Task{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	mountRouter(r, "/users", routes.UserRouter())
	mountRouter(r, "/tasks", routes.TasksRouter())

	log.Println("Server running on PORT " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":3001", r))
}
