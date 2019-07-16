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
	udi, err := buildUserData(c, appConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	ruc := registration.RegistrationUseCase{UserDataInterface: udi}
	c.Put(key, &ruc)

	return &ruc, nil
}





