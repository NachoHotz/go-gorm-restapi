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

  if users == nil || len(users) == 0 {
    w.WriteHeader(http.StatusNotFound)
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{
      "status": 404,
      "message": "Users not found"
    }`))
    return
  }

  json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
  var uniqueUser models.User
  var userId = mux.Vars(r)["id"]

  db.DB.First(&uniqueUser, userId)

  if uniqueUser.ID == 0 {
    w.WriteHeader(http.StatusNotFound)
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{
      "status": 404,
      "message": "User not found"
    }`))
    return
  } else {
    db.DB.Model(&uniqueUser).Association("Tasks").Find(&uniqueUser.Tasks)
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
    return
  }

  json.NewEncoder(w).Encode(&user)
}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
  var userId = mux.Vars(r)["id"]
  var user models.User

  db.DB.First(&user, userId)

  if user.ID == 0 {
    w.WriteHeader(http.StatusNotFound)
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{
      "status": 404,
      "message": "User not found"
    }`))
    return
  } else {
    db.DB.Delete(&user)
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{
      "status": 200,
      "message": "User deleted"
    }`))
  }
}