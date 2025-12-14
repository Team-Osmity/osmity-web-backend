package handler

import "github.com/gin-gonic/gin"

// Health godoc
// @Summary Health check
// @Tags    system
// @Produce json
// @Success 200 {object} handler.HealthResponse
// @Router  /api/health [get]
func Health(c *gin.Context) {
    c.JSON(200, HealthResponse{
        Status: "ok",
    })
}
