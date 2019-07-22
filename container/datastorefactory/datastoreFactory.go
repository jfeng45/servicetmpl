// Package datastorefactory using factory method pattern to create concrete database handler.
// Datastore can be a database or a service ( for example, gRPC service or RESTFul service), which provides data access
// for domain model.
// There is only one method Build() for the factory and all different types of store following the same interface
// to build the store connection.
// Generally speaking, each data store need a separate factory.
package datastorefactory

import (
	"github.com/jfeng45/servicetmpl/config"
	"github.com/jfeng45/servicetmpl/container"
)


// To map "database code" to "database interface builder"
// Concreate builder is in corresponding factory file. For example, "sqlFactory" is in "sqlFactory".go
var dsFbMap = map[string]dsFbInterface{
	config.SQLDB:      &sqlFactory{},
	config.COUCHDB:    &couchdbFactory{},
	config.CACHE_GRPC: &cacheGrpcFactory{},
}

// DataStoreInterface serve as a marker to indicate the return type for Build method
type DataStoreInterface interface{}

// The builder interface for factory method pattern
// Every factory needs to implement Build method
type dsFbInterface interface {
	Build(container.Container, *config.DataStoreConfig) (DataStoreInterface, error)
}

//GetDataStoreFb is accessors for factoryBuilderMap
func GetDataStoreFb(key string) dsFbInterface {
	return dsFbMap[key]
}
