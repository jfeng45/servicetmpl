package loggerfactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container/loggerfactory/zap"
	"github.com/pkg/errors"
)

// receiver for zap factory
type ZapFactory struct {}

// build zap logger
func (mf *ZapFactory) Build(lc *configs.LogConfig) error {
	key := lc.Code
	if ZAP != lc.Code {
		errMsg := ZAP + " in zapFactory doesn't match key = " + key
		return errors.New(errMsg)
	}
	err := zap.RegisterLog(*lc)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}
