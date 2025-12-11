package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"

    // Swagger docs
    _ "osmity-web-backend/docs"
    httpSwagger "github.com/swaggo/http-swagger"

    // ハンドラ層
    "osmity-web-backend/handler"
)

// @title Osmity Backend API
// @version 1.0
// @description API for osmity.com and shizuku86.com services
// @host localhost:8080
// @BasePath /api
func main() {

    // -----------------------------
    // PORT の読み込み
    // -----------------------------
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // デフォルト
    }

    r := mux.NewRouter()

    // Swagger UI
    r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

    // API root
    api := r.PathPrefix("/api").Subrouter()

    api.HandleFunc("/hello", handler.HelloHandler).Methods("GET")
    api.HandleFunc("/blogs", handler.GetBlogsHandler).Methods("GET")
    api.HandleFunc("/blogs", handler.CreateBlogHandler).Methods("POST")

    log.Printf("Starting server on :%s...\n", port)
    if err := http.ListenAndServe(":"+port, r); err != nil {
        log.Fatal(err)
    }
}
