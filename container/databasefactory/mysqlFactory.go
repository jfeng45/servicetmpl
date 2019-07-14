package databasefactory

import (
	"database/sql"
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/jfeng45/servicetmpl/dataservice/gdbc"
	"github.com/jfeng45/servicetmpl/dataservice/userdata/mysql"
	"github.com/pkg/errors"
)

// mysqlFactory is receiver for Build method
//type mysqlFactory dbFactoryBuilder
type mysqlFactory struct {}

// implement Build method for MySQL database
func (mf *mysqlFactory) Build(c container.Container, dc *configs.DatabaseConfig) (dataservice.UserDataInterface, error) {
	logger.Log.Debug("mySqlFactory")
	key := dc.Code
	db, err := sql.Open(dc.DriverName, dc.DataSourceName)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	// check the connection
	err =db.Ping()
	if err != nil {
		return nil,  errors.Wrap(err, "")
	}
	dt := gdbc.DBTx{DB: db}
	uts := mysql.DBTxStore{DB: &dt}
	logger.Log.Debug("uts", uts)
	c.Put(key, &uts)
	return &uts, nil

}

