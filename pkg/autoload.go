package pkg

import (
	"github.com/masraga/meraki/pkg/app"
	"github.com/masraga/meraki/pkg/driver"
	"github.com/masraga/meraki/pkg/helpers"
)

type Autoload struct {
	Config *Config
}

/*
DRIVERS
*/
func (auto *Autoload) Database() *driver.MongodbDriver {
	return driver.NewMongodbdriver(auto.Config.DbUri, auto.Config.DbName)
}

/*
APP
*/
func (auto *Autoload) MongoRepository(CollName string) *app.MongoRepository {
	return &app.MongoRepository{
		CollName: CollName,
		Db:       auto.Database(),
	}
}

/*
HELPERS
*/
func (auto *Autoload) FilenameFormatHelper(file string, extension string) string {
	return helpers.FilenameFormatHelper(file, extension)
}

func (auto *Autoload) JwtHelper() *helpers.JwtHelper {
	return helpers.NewJwtHelper(auto.Config.JwtSecret)
}

func NewAutoload() *Autoload {
	config := *NewConfig("./")
	return &Autoload{
		Config: &config,
	}
}
