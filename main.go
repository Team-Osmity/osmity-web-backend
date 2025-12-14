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

var (
    AppEnv    = "unknown"
    Version   = "unknown"
    BuildTime = "unknown"
    GitCommit = "unknown"
)


func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Starting backend | APP_ENV=%s | PORT=%s", AppEnv, port)

    r := mux.NewRouter()

    
    // Swagger は dev のみ
    if AppEnv == "dev" {
        r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
    }

    api := r.PathPrefix("/api").Subrouter()
    api.HandleFunc("/hello", handler.HelloHandler).Methods("GET")
    api.HandleFunc("/blogs", handler.GetBlogsHandler).Methods("GET")
    api.HandleFunc("/blogs", handler.CreateBlogHandler).Methods("POST")

    log.Fatal(http.ListenAndServe(":"+port, r))
}
