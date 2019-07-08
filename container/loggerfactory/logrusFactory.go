package logger

import (
	"github.com/jfeng45/servicetmpl/appcontainer/loggerfactory/logrus"
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/pkg/errors"
)

// receiver for logrus factory
type logrusFactory logFactoryBuilder

// build logrus logger
func (mf *logrusFactory) Build(lc *configs.LogConfig) error {
	err := logrus.RegisterLog(*lc)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}
