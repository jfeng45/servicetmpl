// package databasefactory using factory method pattern to create concrete database handler
package databasefactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/container/registry"
	"github.com/jfeng45/servicetmpl/dataservice"
)

// Empty struct to server as a receiver for build method
type dbFactoryBuilder struct {}

// The builder interface for factory method pattern
// Every factory needs to implement build method
type dbFbInterface interface {
	build(map[string]interface{}, *configs.DatabaseConfig) ( dataservice.UserDataInterface, error)
}

// database code. Need to map to the database code in the configuration yaml file.
const (
	MYSQL   string ="mysql"
	COUCHDB string ="couch"

)

// builder map to map database code to database interface builder
// Concreate builder is in corresponding factory file. For example, "mysqlFactory" is in "mysqlFactory".go
var dbFactoryBuilderMap = map[string]dbFbInterface {
	MYSQL: &mysqlFactory{},
	COUCHDB: &couchdbFactory{},
}

//GetDbFactoryBuilder is accessors for factoryBuilderMap
func GetDbFactoryBuilder(key string) dbFbInterface {
	return dbFactoryBuilderMap[key]
}

// RetrieveUserData retrieves the UserDataInterface from registry.
// If UserDataInterface is not in registry, it will call corresponding builder factory to created UserDataInterface
// based on database code
func RetrieveUserData(factoryMap map[string]interface{}, dc *configs.DatabaseConfig, dfbi dbFbInterface)  (dataservice.UserDataInterface, error){
	key := dc.Code
	logger.Log.Debug("RetrieveUserData: dbc.driverName=", key)
	value, found := registry.GetFromRegistry(factoryMap, key)
	if found {
		logger.Log.Debug("found RetrieveListUser: key=", key)
		return value.(dataservice.UserDataInterface), nil
	}
	// not in map, need to create one
	logger.Log.Debugf("doesn't find key=%v need to craeted a new one\n", key)
	return  dfbi.build(factoryMap, dc)
}




