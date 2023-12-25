package pkg

import "github.com/spf13/viper"

type Config struct {
	DbName string
	DbUri  string
}

func NewConfig(path string) (config *Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		return nil
	}

	viper.Unmarshal(&config)
	return
}
