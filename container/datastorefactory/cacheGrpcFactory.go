package datastorefactory

import (
	"github.com/jfeng45/servicetmpl/config"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// cacheGrpcFactory is an empty receiver for Build method
type cacheGrpcFactory struct{}

func (cgf *cacheGrpcFactory) Build(c container.Container, dsc *config.DataStoreConfig) (DataStoreInterface, error) {
	key := dsc.Code
	//if it is already in container, return
	if value, found := c.Get(key); found {
		logger.Log.Debug("find CacheGrpc key=%v \n", key)
		return value.(*grpc.ClientConn), nil
	}
	//not in map, need to create one
	logger.Log.Debug("doesn't find cacheGrpc key=%v need to created a new one\n", key)

	conn, err := grpc.Dial(dsc.UrlAddress, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	c.Put(key, conn)
	return conn, err
}
