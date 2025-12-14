package main

import (
    "github.com/gin-gonic/gin"
    "osmity-web-backend/internal/router"

	swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"

    _ "osmity-web-backend/docs"
)

func main() {
    r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    router.Register(r)
    r.Run(":8080")
}