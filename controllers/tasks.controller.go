package controllers

import (
  "encoding/json"
  "net/http"

  "github.com/gorilla/mux"
  "github.com/nachohotz/go-gorm-restapi/db"
  "github.com/nachohotz/go-gorm-restapi/db/models"
)

var contentTypeTask = "application/json"

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
  var tasks []models.Task

  db.DB.Find(&tasks)

  if tasks == nil || len(tasks) == 0 {
    w.WriteHeader(http.StatusNotFound)
    w.Header().Set("Content-Type", contentTypeTask)
    w.Write([]byte(`{
      "status": 404,
      "message": "Tasks not found"
      }`))
    return
  }

  json.NewEncoder(w).Encode(tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
  var taskID = mux.Vars(r)["id"]
  var task models.Task

  db.DB.First(&task, taskID)

  if task.ID == 0 {
    w.WriteHeader(http.StatusNotFound)
    w.Header().Set("Content-Type", contentTypeTask)
    w.Write([]byte(`{
      "status": 404,
      "message": "Task not found"
      }`))
    return
  }

  json.NewEncoder(w).Encode(task)
}

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
  var task models.Task

  json.NewDecoder(r.Body).Decode(&task)
  createdTask := db.DB.Create(&task)
  err := createdTask.Error

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(err.Error()))
    return
  }

  json.NewEncoder(w).Encode(&task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
  var taskID = mux.Vars(r)["id"]
  var task models.Task

  db.DB.First(&task, taskID)

  if task.ID == 0 {
    w.WriteHeader(http.StatusNotFound)
    w.Header().Set("Content-Type", contentTypeTask)
    w.Write([]byte(`{
      "status": 404,
      "message": "Task not found"
      }`))
    return
  }
  db.DB.Delete(&task)

  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", contentTypeTask)
  w.Write([]byte(`{
    "status": 200,
    "message": "Task deleted"
    }`))
}
