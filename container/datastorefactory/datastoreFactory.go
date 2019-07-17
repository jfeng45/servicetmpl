// Package datastorefactory using factory method pattern to create concrete database handler.
// Datastore can be a database or a service ( for example, gRPC service or RESTFul service), which provides data access
// for domain model.
// There is only one method Build() for the factory and all different types of store following the same interface
// to build the store connection.
// Generally speaking, each data store need a separate factory.
package datastorefactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
)

// database code. Need to map to the database code in the configuration yaml file.
const (
	SQL        string ="sql"
	COUCHDB    string ="couch"
	CACHE_GRPC string = "cacheGrpc"
)
// To map "database code" to "database interface builder"
// Concreate builder is in corresponding factory file. For example, "sqlFactory" is in "sqlFactory".go
var dbFbMap = map[string]dbFbInterface {
	SQL:        &sqlFactory{},
	COUCHDB:    &couchdbFactory{},
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





