package datastorefactory

import (
	"database/sql"
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/tools/gdbc"
	"github.com/pkg/errors"
)

// mysqlFactory is receiver for Build method
type mysqlFactory struct {}

// implement Build method for MySQL database
func (mf *mysqlFactory) Build(c container.Container, dc *configs.DataStoreConfig) (DataStoreInterface, error) {
	logger.Log.Debug("mySqlFactory")
	key := dc.Code
	db, err := sql.Open(dc.DriverName, dc.UrlAddress)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	// check the connection
	err =db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	dt := gdbc.DBTx{DB: db}
	c.Put(key, &dt)
	return &dt, nil

}

