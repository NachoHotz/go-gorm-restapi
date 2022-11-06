package routes

import (
  "github.com/gorilla/mux"
  "github.com/nachohotz/go-gorm-restapi/controllers"
)

func TasksRouter() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)

  router.HandleFunc("/", controllers.GetTasksHandler).Methods("GET")
  router.HandleFunc("/{id}", controllers.GetTaskHandler).Methods("GET")
  router.HandleFunc("/", controllers.PostTaskHandler).Methods("POST")
  router.HandleFunc("/{id}", controllers.DeleteTaskHandler).Methods("DELETE")

  return router
}
