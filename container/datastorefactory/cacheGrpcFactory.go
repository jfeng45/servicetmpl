package datastorefactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// cacheGrpcFactory is receiver for Build method
type cacheGrpcFactory struct {}

func (cgfb *cacheGrpcFactory) Build(c container.Container, dsc *configs.DataStoreConfig) (DataStoreInterface, error) {
	key := dsc.Code

	if CACHE_GRPC != key {
		errMsg := CACHE_GRPC + " in cacheGrpcFactory doesn't match key = " + key
		return nil, errors.New(errMsg)
	}

	value, found := c.Get(key)
	if found {
		logger.Log.Debug("find CacheGrpc key=%v \n",key)
		return value.(dataservice.CacheDataInterface), nil
	}
	//not in map, need to create one
	logger.Log.Debug("doesn't find cacheGrpc key=%v need to created a new one\n",key)

	conn, err := grpc.Dial(dsc.UrlAddress, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	c.Put(key, conn)
	return conn, err
}



