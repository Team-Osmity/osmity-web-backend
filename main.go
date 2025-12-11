package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/api/kon", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "konnitiwa!")
    })
    
    fmt.Println("Backend running on :8080")
    http.ListenAndServe(":8080", nil)
}
