package main

import (
	"osmity-web-backend/internal/buildinfo"
	"osmity-web-backend/internal/router"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "osmity-web-backend/docs"
)

func main() {
    r := gin.Default()
    if buildinfo.AppEnv == "dev" {
        r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    }
    router.Register(r)
    r.Run(":8080")
}
