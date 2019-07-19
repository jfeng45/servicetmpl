// Package config reasd configurations from a YAML file and load them into a AppConfig type to save the configuration
// information for the application.
// Configuration for different environment can be saved in files with different suffix, for example [Dev], [Prod]
package configs

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// AppConfig represents the application config
type AppConfig struct {
	SQLConfig     DataStoreConfig `yaml:"sqlConfig"`
	CouchdbConfig   DataStoreConfig `yaml:"couchdbConfig"`
	CacheGrpcConfig DataStoreConfig      `yaml:"cacheGrpcConfig"`
	UserGrpcConfig DataStoreConfig      `yaml:"userGrpcConfig"`
	ZapConfig       LogConfig       `yaml:"zapConfig"`
	LorusConfig     LogConfig       `yaml:"logrusConfig"`
	Log             LogConfig       `yaml:"logConfig"`
	UseCase         UseCaseConfig   `yaml:"useCaseConfig"`
}

// UseCaseConfig represents different use cases
type UseCaseConfig struct {
	Registration    RegistrationConfig `yaml:"registration"`
	ListUser    ListUserConfig `yaml:"listUser"`
	ListCourse    ListCourseConfig `yaml:"listCourse"`
}

// RegistrationConfig represents registration use case
type RegistrationConfig struct {
	Code           string          `yaml:"code"`
	UserDataConfig DataConfig    `yaml:"userDataConfig"`
	TxDataConfig       DataConfig `yaml:"txDataConfig"`
}

// ListUserConfig represents list user use case
type ListUserConfig struct {
	Code           string         `yaml:"code"`
	UserDataConfig DataConfig `yaml:"userDataConfig"`
	CacheDataConfig    DataConfig     `yaml:"cacheDataConfig"`
}

// ListCourseConfig represents list course use case
type ListCourseConfig struct {
	Code           string         `yaml:"code"`
	CourseDataConfig DataConfig `yaml:"courseDataConfig"`
}

// DataConfig represents data service
type DataConfig struct {
	Code            string          `yaml:"code"`
	DataStoreConfig DataStoreConfig `yaml:"dataStoreConfig"`
}

// DataConfig represents handlers for data store. It can be a database or a gRPC connection
type DataStoreConfig struct{
	Code        string         `yaml:"code"`
	// Only database has a driver name, for grpc it is "tcp" ( network) for server
	DriverName string `yaml:"driverName"`
	// For database, this is datasource name; for grpc, it is target url
	UrlAddress string `yaml:"urlAddress"`
	// Only some databases need this database name
	DbName string `yaml:"dbName"`
}

// LogConfig represents logger handler
// Logger has many parameters can be set or changed. Currently, only three are listed here. Can add more into it to
// fits your needs.
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