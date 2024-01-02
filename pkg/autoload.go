package pkg

import (
	"github.com/masraga/meraki/pkg/driver"
)

type Autoload struct {
	Config *Config
}

func (auto *Autoload) Database() *driver.MongodbDriver {
	return driver.NewMongodbdriver(auto.Config.DbUri, auto.Config.DbName)
}

func NewAutoload() *Autoload {
	config := *NewConfig("./")
	return &Autoload{
		Config: &config,
	}
}
