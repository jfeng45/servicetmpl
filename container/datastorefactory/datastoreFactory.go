// package datastorefactory using factory method pattern to create concrete database handler
package datastorefactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
)

// Empty struct to server as a receiver for Build method
//type dbFactoryBuilder struct {}

type DataStoreInterface interface{}
// The builder interface for factory method pattern
// Every factory needs to implement Build method
type dbFbInterface interface {
	Build(container.Container, *configs.DatabaseConfig) (DataStoreInterface, error)
}
// database code. Need to map to the database code in the configuration yaml file.
const (
	MYSQL   string ="mysql"
	COUCHDB string ="couch"
)
// builder map to map database code to database interface builder
// Concreate builder is in corresponding factory file. For example, "mysqlFactory" is in "mysqlFactory".go
var dbFbMap = map[string]dbFbInterface {
	MYSQL: &mysqlFactory{},
	COUCHDB: &couchdbFactory{},
}
//GetDataStoreFb is accessors for factoryBuilderMap
func GetDataStoreFb(key string) dbFbInterface {
	return dbFbMap[key]
}





