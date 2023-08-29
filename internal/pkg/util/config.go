package util

import (
	"time"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App     `mapstructure:"app"`
		Redis   `mapstructure:"redis"`
		Dataset `mapstructure:"dataset"`
	}
	App struct {
		Port string `mapstructure:"port"`
	}
	Redis struct {
		Address  string        `mapstructure:"address"`
		Password string        `mapstructure:"password"`
		Name     int           `mapstructure:"name"`
		Timeout  time.Duration `mapstructure:"timeout"`
	}
	Dataset struct {
		Path string `mapstructure:"path"`
	}
)

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("local")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
