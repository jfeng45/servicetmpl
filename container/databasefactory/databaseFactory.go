// package databasefactory using factory method pattern to create concrete database handler
package databasefactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/dataservice"
)

// Empty struct to server as a receiver for Build method
//type dbFactoryBuilder struct {}

// The builder interface for factory method pattern
// Every factory needs to implement Build method
type dbFbInterface interface {
	Build(container.Container, *configs.DatabaseConfig) ( dataservice.UserDataInterface, error)
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





