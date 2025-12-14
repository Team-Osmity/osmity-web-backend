package main

import (
    "github.com/gin-gonic/gin"
    "osmity-web-backend/internal/router"
)

var (
    AppEnv    = "unknown"
    Version   = "unknown"
    BuildTime = "unknown"
    GitCommit = "unknown"
)

func main() {
    r := gin.Default()
    router.Register(r)
    r.Run(":8080")
}