package dataservicefactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
)

// database code. Need to map to the database code in the configuration yaml file.
const (
	USER_DATA   string ="userData"
	CACHE_DATA   string ="cacheData"
	COURSE_DATA string ="courseData"
)
// builder map to map model data code to model data service interface builder
// each model data service need a separate build
// Concrete builder is in corresponding factory file. For example, "userDataServiceFactory" is in "userDataServiceFactory".go
var dsFbMap = map[string]dataServiceFbInterface {
	USER_DATA: &userDataServiceFactory{},
	CACHE_DATA: &cacheDataServiceFactory{},
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

