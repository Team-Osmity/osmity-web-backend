package handler

import "github.com/gin-gonic/gin"

// Version godoc
// @Summary     Get backend version
// @Description Returns backend build information
// @Tags        system
// @Produce     json
// @Success     200 {object} handler.VersionResponse
// @Router      /api/version [get]
func Version(c *gin.Context) {
    c.JSON(200, VersionResponse{
        Service:   "backend",
        Env:       AppEnv,
        Version:   Version,
        Commit:    GitCommit,
        BuildTime: BuildTime,
    })
}
