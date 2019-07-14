package zap

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/pkg/errors"
)

// receiver for zap factory
type ZapFactory struct {}

// build zap logger
func (mf *ZapFactory) Build(lc *configs.LogConfig) error {

	err := RegisterLog(*lc)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}
