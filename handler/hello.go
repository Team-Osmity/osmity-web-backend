package handler

import (
    "fmt"
    "net/http"
)

// HelloHandler godoc
// @Summary Hello world API
// @Description Returns hello message
// @Tags Example
// @Success 200 {string} string "hello message"
// @Router /hello [get]
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from Osmity backend API! Updated!")
}
