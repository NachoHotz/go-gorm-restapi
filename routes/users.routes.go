package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nachohotz/go-gorm-restapi/db"
	"github.com/nachohotz/go-gorm-restapi/db/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
  var users []models.User
  db.DB.Find(&users)

  json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
  var uniqueUser models.User
  var userId = mux.Vars(r)["id"]

  db.DB.First(&uniqueUser, userId)

  if uniqueUser.ID == 0 {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("User not found"))
  } else {
    json.NewEncoder(w).Encode(&uniqueUser)
  }
}

func PostUsersHandler(w http.ResponseWriter, r *http.Request) {
  var body = r.Body
  var user models.User

  json.NewDecoder(body).Decode(&user)

  createdUser := db.DB.Create(&user)
  err := createdUser.Error

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(err.Error()))
  }

  json.NewEncoder(w).Encode(&user)
}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
  var userId = mux.Vars(r)["id"]
  var user models.User

  db.DB.First(&user, userId)
  db.DB.Delete(&user)

  json.NewEncoder(w).Encode(&user)
}
