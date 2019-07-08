package usecasefactory

import (
	"github.com/jfeng45/servicetmpl/appcontainer/databasefactory"
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/usecase"
	"github.com/jfeng45/servicetmpl/usecase/registration"
	"github.com/pkg/errors"
)

// AddRegistration creates RegistrationUseCaseInterface
func AddRegistration(factoryMap map[string]interface{}, appConfig *configs.AppConfig, key string) (usecase.RegistrationUseCaseInterface, error) {

	uc := &appConfig.UseCase.Registration.UserConfig
	udi, err := databasefactory.RetrieveUserData(factoryMap, uc, databasefactory.GetDbFactoryBuilder(uc.Code))
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	ruc := registration.RegistrationUseCase{UserDataInterface: udi}

	factoryMap[key] = &ruc

	return &ruc, nil
}




