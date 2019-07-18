// package loggerfactory handles creating concrete logger with factory method pattern
package loggerfactory

import (
	"github.com/jfeng45/servicetmpl/configs"
)

// constant for logger code, it needs to match log code (logConfig)in configuration
const (
	LOGRUS string ="logrus"
	ZAP string ="zap"
)
// logger mapp to map logger code to logger builder
var logfactoryBuilderMap = map[string]logFbInterface{
	ZAP: &ZapFactory{},
	LOGRUS: &LogrusFactory{},
}

// interface for logger factory
type logFbInterface interface {
	Build(*configs.LogConfig) error
}

// accessors for factoryBuilderMap
func GetLogFactoryBuilder(key string) logFbInterface {
	return logfactoryBuilderMap[key]
}
