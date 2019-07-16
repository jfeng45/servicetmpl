package dataservicefactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/datastorefactory"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/jfeng45/servicetmpl/adapter/cacheclient"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)
// userDataServiceFactory is a empty receiver for Build method
type cacheDataServiceFactory struct {}

func (udmf *cacheDataServiceFactory) Build(c container.Container, dataConfig *configs.DataConfig) (DataServiceInterface, error) {
	logger.Log.Debug("userDataServiceFactory")
	key := dataConfig.Code
	if CACHE_DATA != key {
		errMsg := USER_DATA + " in userDataServiceFactory doesn't match key = " + key
		return nil, errors.New(errMsg)
	}
	//if it is already in container, return
	if value, found := c.Get(key); found {
		return value.(dataservice.CacheDataInterface), nil
	}
	dsc := dataConfig.DataStoreConfig
	dsi, err := datastorefactory.GetDataStoreFb(dsc.Code).Build(c, &dsc)
	grpcConn := dsi.(*grpc.ClientConn)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	cdg := cacheclient.CacheDataGrpc{grpcConn}
	//logger.Log.Debug("udm:", udm.DB)
	c.Put(key, &cdg)
	return &cdg, nil

	return nil, nil
}

