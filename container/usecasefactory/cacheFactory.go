package usecasefactory

import (
	"github.com/jfeng45/servicetmpl/adapter/cacheclient"
	"github.com/jfeng45/servicetmpl/container/registry"
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/jfeng45/servicetmpl/tools/logger"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// addCacheData creates CacheDataInterface
func addCacheData(factoryMap map[string]interface{}, gc *configs.GrpcConfig) (dataservice.CacheDataInterface, error) {
	conn, err := grpc.Dial(gc.Target, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	cdg := cacheclient.CacheDataGrpc{*conn}
	key := gc.Target
	factoryMap[key] = &cdg
	return &cdg, err
}

// retrieveCache retrieves CacheDataInterface from factory map. If it is not in map yet, it created one and put it
// into map.
// CacheDataInterface is a singleton
func retrieveCache(factoryMap map[string]interface{}, gc *configs.GrpcConfig) (dataservice.CacheDataInterface, error){
	key := gc.Code
	value, found := registry.GetFromRegistry(factoryMap, key)
	if found {
		logger.Log.Debug("find RetrieveCache key=%v \n",key)
		return value.(dataservice.CacheDataInterface), nil
	}
	//not in map, need to create one
	logger.Log.Debug("doesn't find key=%v need to created a new one\n",key)
	return addCacheData(factoryMap, gc)
}


