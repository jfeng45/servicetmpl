package usecasefactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/dataservicefactory"
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/pkg/errors"
)

func buildUserData (c container.Container, appConfig *configs.AppConfig, ) (dataservice.UserDataInterface, error){
	uc := &appConfig.UseCase.Registration.UserDataConfig
	dsi, err := dataservicefactory.GetDataServiceFb(uc.Code).Build(c, uc )
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	udi := dsi.(dataservice.UserDataInterface)
	return udi, nil
}

