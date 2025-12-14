package handler

import (
    "fmt"
    "net/http"
    "os" // ← これを追加
)

// HelloHandler godoc
// @Summary Hello world API
// @Description Returns hello message
// @Tags Example
// @Success 200 {string} string "hello message"
// @Router /hello [get]
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    env := os.Getenv("APP_ENV")
    if env == "" {
        env = "unknown"
    }

    fmt.Fprintf(
        w,
        "Hello from Osmity backend API! Updated! (env=%s)",
        env,
    )
}
