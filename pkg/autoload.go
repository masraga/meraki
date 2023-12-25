package pkg

import (
	"github.com/masraga/meraki/pkg/driver"
)

type Autolaod struct {
	Config *Config
}

func (auto *Autolaod) Database() *driver.MongodbDriver {
	return driver.NewMongodbdriver(auto.Config.DbUri, auto.Config.DbName)
}

func NewAutoload() *Autolaod {
	config := *NewConfig("./")
	return &Autolaod{
		Config: &config,
	}
}
