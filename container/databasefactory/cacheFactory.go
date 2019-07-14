package databasefactory

import (
	"github.com/jfeng45/servicetmpl/adapter/cacheclient"
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Build creates concrete type for CacheDataInterface
//func Build(c container.Container, gc *configs.GrpcConfig) (dataservice.CacheDataInterface, error) {
//	conn, err := grpc.Dial(gc.Target, grpc.WithInsecure())
//	if err != nil {
//		return nil, errors.Wrap(err, "")
//	}
//
//	cdg := cacheclient.CacheDataGrpc{*conn}
//	key := gc.Target
//	c.Put(key, &cdg)
//	return &cdg, err
//}

func Build(c container.Container, gc *configs.GrpcConfig) (dataservice.CacheDataInterface, error) {
	key := gc.Code
	//value, found := registry.GetFromRegistry(factoryMap, key)
	value, found := c.Get(key)
	if found {
		logger.Log.Debug("find RetrieveCache key=%v \n",key)
		return value.(dataservice.CacheDataInterface), nil
	}
	//not in map, need to create one
	logger.Log.Debug("doesn't find cache key=%v need to created a new one\n",key)

	conn, err := grpc.Dial(gc.Target, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	cdg := cacheclient.CacheDataGrpc{*conn}
	c.Put(key, &cdg)
	return &cdg, err
}

// retrieveCache retrieves CacheDataInterface from factory map. If it is not in map yet, it created one and put it
// into map.
// CacheDataInterface is a singleton
//func retrieveCache(c container.Container, gc *configs.GrpcConfig) (dataservice.CacheDataInterface, error){
//	key := gc.Code
//	//value, found := registry.GetFromRegistry(factoryMap, key)
//	value, found := c.Get(key)
//	if found {
//		logger.Log.Debug("find RetrieveCache key=%v \n",key)
//		return value.(dataservice.CacheDataInterface), nil
//	}
//	//not in map, need to create one
//	logger.Log.Debug("doesn't find key=%v need to created a new one\n",key)
//	return Build(c, gc)
//}


