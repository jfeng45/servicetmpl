package userdataservicefactory


import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/datastorefactory"
	"github.com/jfeng45/servicetmpl/dataservice"
)

var udsFbMap = map[string] userDataServiceFbInterface {
	datastorefactory.SQLDB:   &sqlUserDataServiceFactory{},
	datastorefactory.COUCHDB: &couchdbUserDataServiceFactory{},
}

// The builder interface for factory method pattern
// Every factory needs to implement Build method
type userDataServiceFbInterface interface {
	Build(container.Container, *configs.DataConfig) (dataservice.UserDataInterface, error)
}

// GetDataServiceFb is accessors for factoryBuilderMap
func GetUserDataServiceFb(key string) userDataServiceFbInterface {
	return udsFbMap[key]
}




