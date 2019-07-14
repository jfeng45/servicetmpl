package logrus

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/pkg/errors"
)

// receiver for logrus factory
type LogrusFactory struct {}

// build logrus logger
func (mf *LogrusFactory) Build(lc *configs.LogConfig) error {
	err := RegisterLog(*lc)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}
