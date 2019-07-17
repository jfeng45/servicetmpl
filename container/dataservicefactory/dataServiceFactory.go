// package dataservicefactory using factory method pattern to create concrete type to provide persistence service
// the source of data can come from database ( for domain model "user") or from other service ( for domain model cache,
// it is a gRPC service).
// There is only one method Build() for the factory and all different types of data service following the same interface
// to build the data service.

package dataservicefactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
)

// database code. Need to map to the database code in the configuration yaml file.
const (
	USER_DATA   string ="userData"
	CACHE_DATA   string ="cacheData"
	TX_DATA   string ="txData"
	COURSE_DATA string ="courseData"
)
// To map "model data code" to "model data service interface builder"
// each model data service need a separate build
// Concrete builder is in corresponding factory file. For example, "userDataServiceFactory" is in "userDataServiceFactory".go
var dsFbMap = map[string]dataServiceFbInterface {
	USER_DATA: &userDataServiceFactory{},
	CACHE_DATA: &cacheDataServiceFactory{},
	TX_DATA: &txDataServiceFactory{},
	//COURSE: &courseDataServiceFactory{},
}

// DataServiceInterface serves as a marker to indicate the return type for Build method
type DataServiceInterface interface{}

// The builder interface for factory method pattern
// Every factory needs to implement Build method
type dataServiceFbInterface interface {
	Build(container.Container, *configs.DataConfig) (DataServiceInterface, error)
}

// GetDataServiceFb is accessors for factoryBuilderMap
func GetDataServiceFb(key string) dataServiceFbInterface {
	return dsFbMap[key]
}

