package router

import (
    "github.com/gin-gonic/gin"
    "osmity-web-backend/internal/handler"
)

func Register(r *gin.Engine) {
    api := r.Group("/api")
    {
        api.GET("/version", handler.Version)
        api.GET("/health", handler.Health)
    }
}