//package container use dependency injection to create concrete type and wire the whole application together
package servicecontainer

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container/logger"
	logFactory "github.com/jfeng45/servicetmpl/container/loggerfactory"
	"github.com/jfeng45/servicetmpl/container/registry"
	"github.com/jfeng45/servicetmpl/container/usecasefactory"
	"github.com/jfeng45/servicetmpl/usecase"
	"github.com/pkg/errors"
)

type ServiceContainer struct {
	FactoryMap map[string]interface{}
	Config     *configs.AppConfig
}

func (sc *ServiceContainer)InitApp( filename string ) error {
	var err error
	sc.Config, err = loadConfig(filename)
	if err != nil {
		return errors.Wrap(err,"loadConfig")
	}
	err = loadLogger(sc.Config.Log)
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
func loadConfig (filename string) (*configs.AppConfig, error) {

	ac, err := configs.ReadConfig(filename)
	if err != nil {
		return nil, errors.Wrap(err, "read container")
	}
	return &ac, nil
}

func (sc *ServiceContainer) RetrieveRegistration() (usecase.RegistrationUseCaseInterface, error){

	key := sc.Config.UseCase.Registration.Code
	value, found := registry.GetFromRegistry(sc.FactoryMap, key)
	if found {
		logger.Log.Debug("found Retrieve registration: key=", key)
		return value.(usecase.RegistrationUseCaseInterface), nil
	}
	//not in map, need to create one
	logger.Log.Debugf("doesn't find key=%v need to created a new one\n",key)
	return usecasefactory.AddRegistration(sc.FactoryMap, sc.Config, key)
}

func (sc *ServiceContainer) RetrieveListUser() (usecase.ListUserUseCaseInterface, error){
	key := sc.Config.UseCase.ListUser.Code
	value, found := registry.GetFromRegistry(sc.FactoryMap, key)
	if found {
		logger.Log.Debug("found RetrieveListUser: key=", key)
		return value.(usecase.ListUserUseCaseInterface), nil
	}
	//not in map, need to create one
	logger.Log.Debugf("doesn't find key=%v need to created a new one\n",key)
	return usecasefactory.AddListUser(sc.FactoryMap, sc.Config, key)
}
