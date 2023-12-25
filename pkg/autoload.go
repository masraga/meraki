package pkg

import "github.com/masraga/meraki/pkg/driver"

type Autolaod struct {
	Config *Config
}

func (auto *Autolaod) Database() *driver.MongodbDriver {
	return driver.NewMongodbDriver(auto.Config.DbUri, auto.Config.DbName)
}

func NewAutoload(config Config) *Autolaod {
	return &Autolaod{
		Config: &config,
	}
}
