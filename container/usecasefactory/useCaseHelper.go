package usecasefactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/dataservicefactory"
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/pkg/errors"
)


func buildUserData (c container.Container, dc *configs.DataConfig) (dataservice.UserDataInterface, error){
	dsi, err := dataservicefactory.GetDataServiceFb(dc.Code).Build(c, dc )
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	udi := dsi.(dataservice.UserDataInterface)
	return udi, nil
}

func buildTxData (c container.Container, dc *configs.DataConfig) (dataservice.TxDataInterface, error){
	dsi, err := dataservicefactory.GetDataServiceFb(dc.Code).Build(c, dc )
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	tdi := dsi.(dataservice.TxDataInterface)
	return tdi, nil
}

func buildCacheData (c container.Container,  dc *configs.DataConfig) (dataservice.CacheDataInterface, error){
	//logger.Log.Debug("uc:", cdc)
	dsi, err := dataservicefactory.GetDataServiceFb(dc.Code).Build(c, dc )
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	cdi := dsi.(dataservice.CacheDataInterface)
	return cdi, nil
}

func buildCourseData (c container.Container, dc *configs.DataConfig) (dataservice.CourseDataInterface, error){
	dsi, err := dataservicefactory.GetDataServiceFb(dc.Code).Build(c, dc )
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	cdi := dsi.(dataservice.CourseDataInterface)
	return cdi, nil
}

