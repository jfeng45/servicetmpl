package loggerfactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container/loggerfactory/logrus"
	"github.com/pkg/errors"
)

// receiver for logrus factory
type LogrusFactory struct {}

// build logrus logger
func (mf *LogrusFactory) Build(lc *configs.LogConfig) error {
	key := lc.Code
	if LOGRUS != lc.Code {
		errMsg := LOGRUS + " in LogrusFactory doesn't match key = " + key
		return errors.New(errMsg)
	}
	err := logrus.RegisterLog(*lc)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}
