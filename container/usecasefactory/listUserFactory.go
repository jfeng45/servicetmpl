package usecasefactory

import (
	"github.com/jfeng45/servicetmpl/container/databasefactory"
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/usecase"
	"github.com/jfeng45/servicetmpl/usecase/listuser"
	"github.com/pkg/errors"
)

// AddListUser creates ListUserUseCaseInterface
func AddListUser(factoryMap map[string]interface{}, appConfig *configs.AppConfig, key string) (usecase.ListUserUseCaseInterface, error) {

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




