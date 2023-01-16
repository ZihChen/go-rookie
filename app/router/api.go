package router

import (
	"github.com/gin-gonic/gin"
	"go-rookie/app/settings"
)

func Setup(r *gin.Engine) {
	if settings.Config.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
}
