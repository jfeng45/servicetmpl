// package logrus handles creating logrus logger
package logrus

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

//type loggerWrapper struct {
//	lw *logrus.Logger
//}
//
//func (logger *loggerWrapper) Errorf(format string, args ...interface{}) {
//	logger.lw.Errorf(format, args)
//}
//func (logger *loggerWrapper) Fatalf(format string, args ...interface{}) {
//	logger.lw.Fatalf(format, args)
//}
//func (logger *loggerWrapper) Fatal(args ...interface{}) {
//	logger.lw.Fatal(args)
//}
//func (logger *loggerWrapper) Infof(format string, args ...interface{}) {
//	logger.lw.Infof(format, args)
//}
//func (logger *loggerWrapper) Warnf(format string, args ...interface{}) {
//	logger.lw.Warnf(format, args)
//}
//func (logger *loggerWrapper) Debugf(format string, args ...interface{}) {
//	logger.lw.Debugf(format, args)
//}
//func (logger *loggerWrapper) Printf(format string, args ...interface{}) {
//	logger.lw.Printf(format, args)
//}
//func (logger *loggerWrapper) Println(args ...interface{}) {
//	logger.lw.Println(args)
//}

func RegisterLog(lc configs.LogConfig) error{
	//standard configuration
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})
	log.SetReportCaller(true)
	//log.SetOutput(os.Stdout)
	//customize it from configuration file
	err := customizeLogFromConfig(log, lc)
	if err != nil {
		return errors.Wrap(err, "")
	}
	//This is for loggerWrapper implementation
	//logger.Logger(&loggerWrapper{log})

	logger.SetLogger(log)
	return nil
}

// customizeLogFromConfig customize log based on parameters from configuration file
func customizeLogFromConfig(log *logrus.Logger, lc configs.LogConfig) error{
	log.SetReportCaller(lc.EnableCaller)
	//log.SetOutput(os.Stdout)
	l := &log.Level
	err := l.UnmarshalText([]byte(lc.Level))
	if err != nil {
		return errors.Wrap(err, "")
	}
	log.SetLevel(*l)
	return nil
}

