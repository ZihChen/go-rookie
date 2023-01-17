package router

import (
	"github.com/gin-gonic/gin"
	"go-rookie/app/logger"
	"go-rookie/app/settings"
)

func Setup(r *gin.Engine) {
	if settings.Config.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r.Use(logger.GinLogger(), logger.GinRecovery(false))
	r.GET("/", func(c *gin.Context) {
		panic(123)
		c.JSON(200, "hello world")
	})
}
