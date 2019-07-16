//package container use dependency injection to create concrete type and wire the whole application together
package servicecontainer

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container/logger"
	logFactory "github.com/jfeng45/servicetmpl/container/loggerfactory"
	"github.com/jfeng45/servicetmpl/container/usecasefactory"
	"github.com/pkg/errors"
)

type ServiceContainer struct {
	FactoryMap map[string]interface{}
}

func (sc *ServiceContainer)InitApp( filename string ) error {
	var err error
	config, err := loadConfig(filename)
	if err != nil {
		return errors.Wrap(err,"loadConfig")
	}
	err = loadLogger(config.Log)
	if err != nil {
		return errors.Wrap(err,"loadLogger")
	}
	err = buildUseCase(sc, config)
	if err != nil {
		return errors.Wrap(err,"build use case")
	}
	return nil
}

func (sc *ServiceContainer) GetInstance(code string) ( interface{}, error) {

	value, found := sc.FactoryMap[code]
	if found {
		logger.Log.Debug("found Retrieve registration: code=", code)
		return value, nil
	} else {
		errMsg := "can't find corresponding type for code " + code + " in container"
		return nil, errors.New(errMsg)
	}
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
func loadConfig (filename string) (*configs.AppConfig, error) {

	ac, err := configs.ReadConfig(filename)
	if err != nil {
		return nil, errors.Wrap(err, "read container")
	}
	return &ac, nil
}

// create concrete types for use case interfaces
func buildUseCase(sc *ServiceContainer, config *configs.AppConfig) error {
	for key,ucfb := range usecasefactory.UseCaseFactoryBuilderMap {
		_, err := ucfb.Build(sc, config, key)
		if err != nil {
			return errors.Wrap(err, "build use case")
		}
	}
	return nil
}

func (sc *ServiceContainer) Get(code string) (interface{}, bool){
	value, found := sc.FactoryMap[code]
	return value, found
}

func (sc *ServiceContainer) Put(code string, value interface{}) {
	sc.FactoryMap[code] =value
}



