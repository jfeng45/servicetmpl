package appcontainer
import (
	"github.com/jfeng45/servicetmpl/appcontainer/databasefactory"
	"github.com/jfeng45/servicetmpl/appcontainer/registry"
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/tools/logger"
	"github.com/jfeng45/servicetmpl/usecase"
	"github.com/jfeng45/servicetmpl/usecase/listuser"
	"github.com/pkg/errors"
)

// listUserFactory creates ListUserUseCaseInterface
func listUserFactory(factoryMap map[string]interface{}, appConfig *configs.AppConfig, key string) (usecase.ListUserUseCaseInterface, error) {

	uc := &appConfig.UseCase.Registration.UserConfig
	udi, err := databasefactory.RetrieveUserData(factoryMap, uc, databasefactory.GetDbFactoryBuilder(uc.Code))
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	cdi, err := retrieveCache(factoryMap, &appConfig.CacheGrpcConfig)
	luuc := listuser.ListUserUseCase{UserDataInterface: udi, CacheDataInterface: cdi}

	factoryMap[key] = &luuc

	return &luuc, nil
}

// RetrieveListUser retrieves ListUserUseCaseInterface from factory map. If it is not in map yet, it created one and put it
// into map.
// RetrieveListUser is a singleton
func RetrieveListUser(factoryMap map[string]interface{}) (usecase.ListUserUseCaseInterface, error){
	config := configs.GetAppConfig()
	key := config.UseCase.ListUser.Code
	value, found := registry.GetFromRegistry(factoryMap, key)
	if found {
		logger.Log.Debug("found RetrieveListUser: key=", key)
		return value.(usecase.ListUserUseCaseInterface), nil
	}
	//not in map, need to create one
	logger.Log.Debugf("doesn't find key=%v need to created a new one\n",key)
	return listUserFactory(factoryMap, config, key)
}

