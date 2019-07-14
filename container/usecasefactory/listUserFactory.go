package usecasefactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/databasefactory"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/jfeng45/servicetmpl/usecase/listuser"
	"github.com/pkg/errors"
)

type ListUserFactory struct {}

func (lf *ListUserFactory)Build(c container.Container, appConfig *configs.AppConfig, key string) error {

	uc := &appConfig.UseCase.Registration.UserConfig
	var udi dataservice.UserDataInterface
	var err error
	value, found := c.Get(uc.Code)
	if found {
		logger.Log.Debug("found ListUser use case: key=", key)
		udi = value.(dataservice.UserDataInterface)
	} else {
		udi, err = databasefactory.GetDbFactoryBuilder(uc.Code).Build(c, uc )
	}

	if err != nil {
		return errors.Wrap(err, "")
	}
	cdi, err := databasefactory.Build(c, &appConfig.CacheGrpcConfig)
	luuc := listuser.ListUserUseCase{UserDataInterface: udi, CacheDataInterface: cdi}

	c.Put(key,&luuc)

	return nil
}






