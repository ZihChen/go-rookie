package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rookie/app/router"
	"go-rookie/app/settings"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var r *gin.Engine

func Run() {
	r = gin.New()

	router.Setup(r)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Config.Port),
		Handler: r,
	}
	zap.L().Debug("server listen", zap.String("port", strconv.Itoa(settings.Config.Port)))

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("listen: %s\n", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
		zap.L().Error("Server Shutdown: ", zap.Error(err))
	}
	zap.L().Info("Server exiting")
}