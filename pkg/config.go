package pkg

import (
	"github.com/spf13/viper"
)

type Config struct {
	DbName string `mapstructure:"DB_NAME"`
	DbUri  string `mapstructure:"DB_URI"`
}

func NewConfig(path string) (config *Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.Unmarshal(&config)
	return
}
