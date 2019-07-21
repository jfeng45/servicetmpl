package usecasefactory

import (
	"github.com/jfeng45/servicetmpl/config"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/usecase/listuser"
	"github.com/pkg/errors"
)

type ListUserFactory struct{}

func (luf *ListUserFactory) Build(c container.Container, appConfig *config.AppConfig, key string) (UseCaseInterface, error) {
	uc := appConfig.UseCase.ListUser

	if container.LIST_USER != uc.Code {
		errMsg := container.LIST_USER + " in ListUserFactory doesn't match key = " + key
		return nil, errors.New(errMsg)
	}

	udi, err := buildUserData(c, &uc.UserDataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	cdi, err := buildCacheData(c, &uc.CacheDataConfig)
	luuc := listuser.ListUserUseCase{UserDataInterface: udi, CacheDataInterface: cdi}
	return &luuc, nil
}
