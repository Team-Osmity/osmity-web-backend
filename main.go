package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"

    _ "osmity-web-backend/docs"
    httpSwagger "github.com/swaggo/http-swagger"

    "osmity-web-backend/handler"
)

func main() {
    // -----------------------------
    // APP_ENV 読み込み
    // -----------------------------
    appEnv := os.Getenv("APP_ENV")
    if appEnv == "" {
        appEnv = "unknown"
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Starting backend | APP_ENV=%s | PORT=%s", appEnv, port)

    r := mux.NewRouter()

    // Swagger は dev のみ
    if appEnv == "dev" {
        r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
    }

    api := r.PathPrefix("/api").Subrouter()
    api.HandleFunc("/hello", handler.HelloHandler).Methods("GET")
    api.HandleFunc("/blogs", handler.GetBlogsHandler).Methods("GET")
    api.HandleFunc("/blogs", handler.CreateBlogHandler).Methods("POST")

    log.Fatal(http.ListenAndServe(":"+port, r))
}
