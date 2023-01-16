package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config *AppConfig

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("viper init failed:", err)
		return err
	}

	if err = viper.Unmarshal(&Config); err != nil {
		fmt.Println("viper Unmarshal err", err)
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file has been modified")
	})

	return err
}

type AppConfig struct {
	Name      string     `mapstructure:"name"`
	Mode      string     `mapstructure:"mode"`
	Port      int        `mapstructure:"port"`
	LogConfig *LogConfig `mapstructure:"log"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
