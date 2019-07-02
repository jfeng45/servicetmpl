package appcontainer
import (
	"github.com/jfeng45/servicetmpl/appcontainer/databasefactory"
	"github.com/jfeng45/servicetmpl/appcontainer/registry"
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/tools"
	"github.com/jfeng45/servicetmpl/usecase"
	"github.com/jfeng45/servicetmpl/usecase/registration"
	"github.com/pkg/errors"
)

// addRegistration creates RegistrationUseCaseInterface
func addRegistration(factoryMap map[string]interface{}, appConfig *configs.AppConfig, key string) (usecase.RegistrationUseCaseInterface, error) {

	uc := &appConfig.UseCase.Registration.UserConfig
	udi, err := databasefactory.RetrieveUserData(factoryMap, uc, databasefactory.GetDbFactoryBuilder(uc.Code))
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	ruc := registration.RegistrationUseCase{UserDataInterface: udi}

	factoryMap[key] = &ruc

	return &ruc, nil
}

// RetrieveRegistration retrieves RegistrationUseCaseInterface from factory map. If it is not in map yet, it created one and put it
// into map.
// RetrieveRegistration is a singleton
func RetrieveRegistration(factoryMap map[string]interface{}) (usecase.RegistrationUseCaseInterface, error){
	config := configs.GetAppConfig()
	key := config.UseCase.Registration.Code
	value, found := registry.GetFromRegistry(factoryMap, key)
	if found {
		tools.Log.Debug("found Retrieve registration: key=", key)
		return value.(usecase.RegistrationUseCaseInterface), nil
	}
	//not in map, need to create one
	tools.Log.Debugf("doesn't find key=%v need to created a new one\n",key)
	return addRegistration(factoryMap, config, key)
}


