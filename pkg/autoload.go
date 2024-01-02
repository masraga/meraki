package pkg

import (
	"github.com/masraga/meraki/pkg/app"
	"github.com/masraga/meraki/pkg/driver"
)

type Autoload struct {
	Config *Config
}

func (auto *Autoload) Database() *driver.MongodbDriver {
	return driver.NewMongodbdriver(auto.Config.DbUri, auto.Config.DbName)
}

func (auto *Autoload) MongoRepository(CollName string) *app.MongoRepository {
	return &app.MongoRepository{
		CollName: CollName,
		Db:       auto.Database(),
	}
}

func NewAutoload() *Autoload {
	config := *NewConfig("./")
	return &Autoload{
		Config: &config,
	}
}
