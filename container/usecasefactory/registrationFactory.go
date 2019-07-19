package usecasefactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/usecase/registration"
	"github.com/pkg/errors"
)

type RegistrationFactory struct {

}
// Build creates concrete type for RegistrationUseCaseInterface
func (rf *RegistrationFactory) Build(c container.Container, appConfig *configs.AppConfig, key string) (UseCaseInterface, error) {
	uc := appConfig.UseCase.Registration

	if container.REGISTRATION != uc.Code {
		errMsg := container.REGISTRATION  + " in RegistrationFactory doesn't match key = " + key
		return nil, errors.New(errMsg)
	}

	udi, err := buildUserData(c, &uc.UserDataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	tdi, err := buildTxData(c, &uc.TxDataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	ruc := registration.RegistrationUseCase{UserDataInterface: udi, TxDataInterface:tdi}

	return &ruc, nil
}





