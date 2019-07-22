package dataservicefactory

import (
	"github.com/jfeng45/servicetmpl/adapter/cacheclient"
	"github.com/jfeng45/servicetmpl/config"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/datastorefactory"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// cacheDataServiceFactory is a empty receiver for Build method
type cacheDataServiceFactory struct{}

func (cdsf *cacheDataServiceFactory) Build(c container.Container, dataConfig *config.DataConfig) (DataServiceInterface, error) {
	logger.Log.Debug("cacheDataServiceFactory")
	dsc := dataConfig.DataStoreConfig
	dsi, err := datastorefactory.GetDataStoreFb(dsc.Code).Build(c, &dsc)
	grpcConn := dsi.(*grpc.ClientConn)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	cdg := cacheclient.CacheDataGrpc{grpcConn}
	//logger.Log.Debug("udm:", udm.DB)

	return &cdg, nil
}
