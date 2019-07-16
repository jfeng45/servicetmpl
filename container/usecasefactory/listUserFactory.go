package usecasefactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/usecase/listuser"
	"github.com/pkg/errors"
)

type ListUserFactory struct {}

func (lf *ListUserFactory)Build(c container.Container, appConfig *configs.AppConfig, key string) (UseCaseInterface, error)  {
	udi, err := buildUserData(c, appConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	cdi, err := buildCacheData(c, appConfig)
	luuc := listuser.ListUserUseCase{UserDataInterface: udi, CacheDataInterface: cdi}
	c.Put(key,&luuc)
	return &luuc, nil
}






