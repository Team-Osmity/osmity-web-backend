package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"

    // Swagger docs
    _ "osmity-web-backend/docs"
    httpSwagger "github.com/swaggo/http-swagger"

    // ハンドラ層（後で作成）
    "osmity-web-backend/handler"
)

// @title Osmity Backend API
// @version 1.0
// @description API for osmity.com and shizuku86.com services
// @host localhost:8080
// @BasePath /api
func main() {
    r := mux.NewRouter()

    // -----------------------------
    // Swagger UI
    // -----------------------------
    r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

    // -----------------------------
    // API ルート
    // -----------------------------
    api := r.PathPrefix("/api").Subrouter()

    // Simple API example
    api.HandleFunc("/hello", handler.HelloHandler).Methods("GET")

    // Blog API base（後で実装）
    api.HandleFunc("/blogs", handler.GetBlogsHandler).Methods("GET")
    api.HandleFunc("/blogs", handler.CreateBlogHandler).Methods("POST")

    // -----------------------------
    // HTTP サーバー起動
    // -----------------------------
    log.Println("Starting server on :8080...")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}
