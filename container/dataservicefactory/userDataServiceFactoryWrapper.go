package dataservicefactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/dataservicefactory/userdataservicefactory"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/pkg/errors"
)

// userDataServiceFactory is a empty receiver for Build method
type userDataServiceFactoryWrapper struct {}

func (udsf *userDataServiceFactoryWrapper) Build(c container.Container, dataConfig *configs.DataConfig) (DataServiceInterface, error) {
	logger.Log.Debug("UserDataServiceFactory")
	userDataKey := dataConfig.Code
	//if it is already in container, return
	if value, found := c.Get(userDataKey); found {
		return value.(dataservice.UserDataInterface), nil
	}

	key := dataConfig.DataStoreConfig.Code
	udsi, err := userdataservicefactory.GetUserDataServiceFb(key).Build(c,dataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return udsi, nil
}

