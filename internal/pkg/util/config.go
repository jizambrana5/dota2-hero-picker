package util

import "github.com/spf13/viper"

/*type Config struct {
	Database `json:"database"`
}*/

type Config struct {
	Address  string `mapstructure:"DB_ADDRESS"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     int    `mapstructure:"DB"`
	Timeout  int    `mapstructure:"DB_TIMEOUT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
