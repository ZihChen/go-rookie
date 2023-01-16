package main

import (
	"fmt"
	"go-rookie/app/logger"
	"go-rookie/app/settings"
	"go.uber.org/zap"
)

func main() {
	// 載入設定檔
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed:%s \n", err)
		return
	}
	// 初始化日誌
	if err := logger.Init(); err != nil {
		fmt.Printf("init settings failed:%s \n", err)
		return
	}
	zap.L().Debug("logger init success!")
	defer zap.L().Sync()
}
