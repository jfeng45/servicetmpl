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

// AppConfig represents the application config
type AppConfig struct {
	MySQLConfig     DatabaseConfig     `yaml:"mySQLConfig"`
	CouchdbConfig   DatabaseConfig     `yaml:"couchdbConfig"`
	CacheGrpcConfig GrpcConfig         `yaml:"cacheGrpcConfig"`
	ZapConfig       LogConfig         `yaml:"zapConfig"`
	LorusConfig       LogConfig         `yaml:"logrusConfig"`
	Log              LogConfig         `yaml:"logConfig"`
	UseCase         UseCaseConfig       `yaml:"useCaseConfig"`
}

// UseCaseConfig represents different use case
type UseCaseConfig struct {
	Registration    RegistrationConfig `yaml:"registration"`
	ListUser    ListUserConfig `yaml:"listUser"`
}

// RegistrationConfig represents registration use case
type RegistrationConfig struct {
	Code           string         `yaml:"code"`
	UserDataConfig UserDataConfig `yaml:"userDataConfig"`
	TxConfig       DatabaseConfig `yaml:"txConfig"`
}

// ListUserConfig represents list user use case
type ListUserConfig struct {
	Code           string         `yaml:"code"`
	UserDataConfig UserDataConfig `yaml:"userDataConfig"`
	CacheConfig    GrpcConfig     `yaml:"cacheConfig"`
}

// UserDataConfig represents user data service
type UserDataConfig struct {
	Code           string         `yaml:"code"`
	DataStoreConfig DatabaseConfig `yaml:"dataStoreConfig"`
}

// DatabaseConfig represents database handler
type DatabaseConfig struct{
	Code        string         `yaml:"code"`
	DriverName string `yaml:"driverName"`
	DataSourceName string `yaml:"dataSourceName"`
	DbName string `yaml:"dbName"`
}

// DatabaseConfig represents gRPC handler
type GrpcConfig struct{
	Code        string   `yaml:"code"`
	GrpcName string `yaml:"grpcName"`
	Target   string `yaml:"target"`
}

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