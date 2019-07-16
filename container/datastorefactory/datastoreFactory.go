// package datastorefactory using factory method pattern to create concrete database handler
package datastorefactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
)

// database code. Need to map to the database code in the configuration yaml file.
const (
	MYSQL   string ="mysql"
	COUCHDB string ="couch"
	CACHE_GRPC string = "cacheGrpc"
)
// builder map to map database code to database interface builder
// Concreate builder is in corresponding factory file. For example, "mysqlFactory" is in "mysqlFactory".go
var dbFbMap = map[string]dbFbInterface {
	MYSQL: &mysqlFactory{},
	COUCHDB: &couchdbFactory{},
	CACHE_GRPC: &cacheGrpcFactory{},
}

// DataStoreInterface serve as a marker to indicate the return type for Build method
type DataStoreInterface interface{}
// The builder interface for factory method pattern
// Every factory needs to implement Build method
type dbFbInterface interface {
	Build(container.Container, *configs.DataStoreConfig) (DataStoreInterface, error)
}

//GetDataStoreFb is accessors for factoryBuilderMap
func GetDataStoreFb(key string) dbFbInterface {
	return dbFbMap[key]
}





