// package logger handles creating concrete logger with factory method pattern
package logger

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container/loggerfactory/logrus"
	"github.com/jfeng45/servicetmpl/container/loggerfactory/zap"
)

// constant for logger code, it needs to match code in configuration
const (
	LOGRUS string ="logrus"
	ZAP string ="zap"
)
// logger mapp to map logger code to logger builder
var logfactoryBuilderMap = map[string]logFbInterface{
	ZAP: &zap.ZapFactory{},
	LOGRUS: &logrus.LogrusFactory{},
}

// interface for logger factory
type logFbInterface interface {
	Build(*configs.LogConfig) error
}

// accessors for factoryBuilderMap
func GetLogFactoryBuilder(key string) logFbInterface {
	return logfactoryBuilderMap[key]
}
