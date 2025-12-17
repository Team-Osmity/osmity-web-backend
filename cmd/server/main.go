package main

import (
	"log"
	"os"
	"osmity-web-backend/internal/db"
	"osmity-web-backend/internal/router"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "osmity-web-backend/docs"
)

func main() {
    db, err := db.Connect()
    if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
    
    r := gin.Default()

    appEnv := os.Getenv("APP_ENV")
    if appEnv == "dev" {
        r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    }
    router.Register(r)

    _ = db
    r.Run(":8080")
}
