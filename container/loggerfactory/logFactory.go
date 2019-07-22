// package loggerfactory handles creating concrete logger with factory method pattern
package loggerfactory

import (
	"github.com/jfeng45/servicetmpl/config"
)

// logger mapp to map logger code to logger builder
var logfactoryBuilderMap = map[string]logFbInterface{
	config.ZAP:    &ZapFactory{},
	config.LOGRUS: &LogrusFactory{},
}

// interface for logger factory
type logFbInterface interface {
	Build(*config.LogConfig) error
}

// accessors for factoryBuilderMap
func GetLogFactoryBuilder(key string) logFbInterface {
	return logfactoryBuilderMap[key]
}
