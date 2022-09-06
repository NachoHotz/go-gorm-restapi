package routes

import (
	"github.com/gorilla/mux"
	"github.com/nachohotz/go-gorm-restapi/controllers"
)

func UserRouter() *mux.Router {
  r := mux.NewRouter()

  r.HandleFunc("/", controllers.GetUsersHandler).Methods("GET")
  r.HandleFunc("/{id}", controllers.GetUserHandler).Methods("GET")
  r.HandleFunc("/", controllers.PostUsersHandler).Methods("POST")
  r.HandleFunc("/{id}", controllers.DeleteUsersHandler).Methods("DELETE")

  return r
}
