// Package routes ...
package routes

import (
  "net/http"
)

// HomeHandler ...
func HomeHandler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello World"))
}
