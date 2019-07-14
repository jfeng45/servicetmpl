package usecasefactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/databasefactory"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/jfeng45/servicetmpl/usecase/registration"
	"github.com/pkg/errors"
)

type RegistrationFactory struct {
}
// Build creates concrete type for RegistrationUseCaseInterface
func (rf *RegistrationFactory) Build(c container.Container, appConfig *configs.AppConfig, key string) error {

	uc := &appConfig.UseCase.Registration.UserConfig
	var udi dataservice.UserDataInterface
	var err error
	value, found := c.Get(uc.Code)
	if found {
		logger.Log.Debug("found registration use case: key=", key)
		udi = value.(dataservice.UserDataInterface)
	} else {
		udi, err = databasefactory.GetDbFactoryBuilder(uc.Code).Build(c, uc )
	}

	if err != nil {
		return errors.Wrap(err, "")
	}
	ruc := registration.RegistrationUseCase{UserDataInterface: udi}

	c.Put(key, &ruc)

	return nil
}




