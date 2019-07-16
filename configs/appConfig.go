// Package config read configurations from a YAML file and load them into a AppConfig type to save the configuration
// information for the application. The configuration information is loaded only once by a control flag.
// configuration for different environment can be saved in files with different suffix, for example [Dev], [Prod]
package configs

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type dataInterface interface{
	getCode() string
	getDataStore() dataStoreInterface
}

type dataStoreInterface interface {
	getCode () string
}
// AppConfig represents the application config
type AppConfig struct {
	MySQLConfig     DataStoreConfig `yaml:"mySQLConfig"`
	CouchdbConfig   DataStoreConfig `yaml:"couchdbConfig"`
	CacheGrpcConfig DataStoreConfig      `yaml:"cacheGrpcConfig"`
	ZapConfig       LogConfig       `yaml:"zapConfig"`
	LorusConfig     LogConfig       `yaml:"logrusConfig"`
	Log             LogConfig       `yaml:"logConfig"`
	UseCase         UseCaseConfig   `yaml:"useCaseConfig"`
}

// UseCaseConfig represents different use case
type UseCaseConfig struct {
	Registration    RegistrationConfig `yaml:"registration"`
	ListUser    ListUserConfig `yaml:"listUser"`
}

// RegistrationConfig represents registration use case
type RegistrationConfig struct {
	Code           string          `yaml:"code"`
	UserDataConfig DataConfig    `yaml:"userDataConfig"`
	TxConfig       DataStoreConfig `yaml:"txConfig"`
}

// ListUserConfig represents list user use case
type ListUserConfig struct {
	Code           string         `yaml:"code"`
	UserDataConfig DataConfig `yaml:"userDataConfig"`
	CacheDataConfig    DataConfig     `yaml:"cacheDataConfig"`
}

// DataConfig represents data service
type DataConfig struct {
	Code            string          `yaml:"code"`
	DataStoreConfig DataStoreConfig `yaml:"dataStoreConfig"`
}

// DataStoreConfig represents database handler
//type DataStoreConfig struct{
//	Code        string         `yaml:"code"`
//	DriverName string `yaml:"driverName"`
//	DataSourceName string `yaml:"dataSourceName"`
//	DbName string `yaml:"dbName"`
//}

type DataStoreConfig struct{
	Code        string         `yaml:"code"`
	// only database has driver name, grpc don't use it
	DriverName string `yaml:"driverName"`
	// for database this is datasource name; for grpc, it is target url
	UrlAddress string `yaml:"urlAddress"`
	// only for some database need this database name
	DbName string `yaml:"DbName"`
}

//// DataStoreConfig represents database handler
//type grpcDataConfig struct{
//	Code        string         `yaml:"code"`
//	GrpcConfig GrpcConfig `yaml:"grpcConfig"`
//}
//
//// DataStoreConfig represents gRPC handler
//type GrpcConfig struct{
//	Code        string   `yaml:"code"`
//	GrpcName string `yaml:"grpcName"`
//	Target   string `yaml:"target"`
//}

// LogConfig represents logger handler
// Logger has many parameters can be set or changed. Currently, only three are listed here, which is most likely to be
// changed at runtime. Can add more into it to fits your needs.
type LogConfig struct{
	// log library name
	Code        string   `yaml:"code"`
	// log level
	Level       string `yaml:"level"`
	// show caller in log message
	EnableCaller   bool `yaml:"enableCaller"`
}

// ReadConfig reads the file of the filename (in the same folder) and put it into the AppConfig
func ReadConfig(filename string) (AppConfig, error) {
	fmt.Println("read from log file: ", filename)
	var ac AppConfig
	file, err :=ioutil.ReadFile(filename)
	if err != nil {
		return ac, errors.Wrap(err, "read error")
	}
	err = yaml.Unmarshal(file, &ac)

	if err != nil {
		return ac, errors.Wrap(err, "unmarshal")
	}
	fmt.Println("appConfig:", ac)
	return ac, nil
}