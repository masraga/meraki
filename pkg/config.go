package pkg

import (
	"github.com/spf13/viper"
)

type Config struct {
	DbName   string `mapstructure:"DB_NAME"`
	DbUri    string `mapstructure:"DB_URI"`
	DbDriver string `mapstructur:"DB_DRIVER"`

	SystemPort string `mapstructure:"SYSTEM_PORT"`
	JwtSecret  string `mapstructure:"JWT_SECRET"`
}

func NewConfig(path string) (config *Config) {
	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.Unmarshal(&config)
	return
}
