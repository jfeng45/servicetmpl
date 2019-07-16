package dataservicefactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
)

// Empty struct to server as a receiver for Build method
//type dataServiceFactory struct {}
type DataServiceInterface interface{}

// The builder interface for factory method pattern
// Every factory needs to implement Build method
type dataServiceFbInterface interface {
	Build(container.Container, *configs.UserDataConfig) (DataServiceInterface, error)
}
// database code. Need to map to the database code in the configuration yaml file.
const (
	USER_DATA   string ="userData"
	COURSE string ="course"
)
// builder map to map database code to database interface builder
// Concreate builder is in corresponding factory file. For example, "mysqlFactory" is in "mysqlFactory".go
var dsFbMap = map[string]dataServiceFbInterface {
	USER_DATA: &userDataServiceFactory{},
	//COURSE: &courseDataServiceFactory{},
}
// GetDataServiceFb is accessors for factoryBuilderMap
func GetDataServiceFb(key string) dataServiceFbInterface {
	return dsFbMap[key]
}

