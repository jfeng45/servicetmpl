package servicecontainer

import (
	"github.com/jfeng45/servicetmpl/config"
	logFactory "github.com/jfeng45/servicetmpl/container/loggerfactory"
	"github.com/jfeng45/servicetmpl/container/usecasefactory"
	"github.com/pkg/errors"
)

type ServiceContainer struct {
	FactoryMap map[string]interface{}
	AppConfig  *config.AppConfig
}

func (sc *ServiceContainer) InitApp(filename string) error {
	var err error
	config, err := loadConfig(filename)
	sc.AppConfig = config
	if err != nil {
		return errors.Wrap(err, "loadConfig")
	}
	err = loadLogger(config.Log)
	if err != nil {
		return errors.Wrap(err, "loadLogger")
	}

	return nil
}

//func (sc *ServiceContainer) GetInstance(code string) ( interface{}, error) {
//
//	value, found := sc.FactoryMap[code]
//	if found {
//		logger.Log.Debug("found instance in container: code=", code)
//		return value, nil
//	} else {
//		errMsg := "can't find corresponding type for code " + code + " in container"
//		return nil, errors.New(errMsg)
//	}
//}

// loads the logger
func loadLogger(lc config.LogConfig) error {
	loggerType := lc.Code
	err := logFactory.GetLogFactoryBuilder(loggerType).Build(&lc)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

// loads the application configurations
func loadConfig(filename string) (*config.AppConfig, error) {

	ac, err := config.ReadConfig(filename)
	if err != nil {
		return nil, errors.Wrap(err, "read container")
	}
	return &ac, nil
}

func (sc *ServiceContainer) BuildUseCase(code string) (interface{}, error) {
	return usecasefactory.GetUseCaseFb(code).Build(sc, sc.AppConfig, code)
}

func (sc *ServiceContainer) Get(code string) (interface{}, bool) {
	value, found := sc.FactoryMap[code]
	return value, found
}

func (sc *ServiceContainer) Put(code string, value interface{}) {
	sc.FactoryMap[code] = value
}
