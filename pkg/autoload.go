package pkg

import (
	"github.com/masraga/meraki/pkg/app"
	"github.com/masraga/meraki/pkg/driver"
	"github.com/masraga/meraki/pkg/utils"
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
	return utils.FilenameFormatHelper(file, extension)
}

func (auto *Autoload) JwtHelper() *utils.JwtHelper {
	return utils.NewJwtHelper(auto.Config.JwtSecret)
}

func NewAutoload() *Autoload {
	config := *NewConfig("./")
	return &Autoload{
		Config: &config,
	}
}
