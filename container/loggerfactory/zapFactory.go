package logger

import (
	"github.com/jfeng45/servicetmpl/container/loggerfactory/zap"
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/pkg/errors"
)

// receiver for logrus factory
type zapFactory logFactoryBuilder

// build zap logger
func (mf *zapFactory) Build(lc *configs.LogConfig) error {

	err := zap.RegisterLog(*lc)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}
