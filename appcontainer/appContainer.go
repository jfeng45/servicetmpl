//package appcontainer use dependency injection to create concrete type and wire the whole application together
package appcontainer

import (
	logFactory "github.com/jfeng45/servicetmpl/appcontainer/loggerfactory"
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/pkg/errors"
)

// InitApp loads the application configurations and logger
func InitApp( factoryMap map[string]interface{},filename string ) error {
	err := loadConfig(filename)
	if err != nil {
		return errors.Wrap(err,"loadConfig")
	}
	config := configs.GetAppConfig()

	err = loadLogger(config.Log)
	if err != nil {
		return errors.Wrap(err,"loadLogger")
	}
	return nil
}

// loads the logger
func loadLogger(lc configs.LogConfig) error {
	loggerType :=lc.Code
	err := logFactory.GetLogFactoryBuilder(loggerType).Build(&lc)
	if err !=nil {
		return errors.Wrap(err, "")
	}
	return nil
}

// loads the application configurations
func loadConfig (filename string) error {

	if !configs.GetReloadConfig() {
		return nil
	}
	ac, err := configs.ReadConfig(filename)
	if err != nil {
		return errors.Wrap(err, "read appcontainer")
	}
	configs.SetAppConfig(&ac)
	//set the flag to false, so won't reload it again
	configs.SetReloadConfig(false)

	return nil
}

