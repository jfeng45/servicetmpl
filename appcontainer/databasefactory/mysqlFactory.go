package databasefactory

import (
	"database/sql"
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/jfeng45/servicetmpl/dataservice/gdbc"
	"github.com/jfeng45/servicetmpl/dataservice/userdata/mysql"
	"github.com/jfeng45/servicetmpl/tools"
	"github.com/pkg/errors"
)

// mysqlFactory is receiver for build method
type mysqlFactory dbFactoryBuilder

// implement build method for MySQL database
func (mf *mysqlFactory) build(factoryMap map[string]interface{}, dc *configs.DatabaseConfig) (dataservice.UserDataInterface, error) {
	tools.Log.Debug("mySqlFactory")
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
	tools.Log.Debug("uts", uts)
	factoryMap[key] = &uts
	return &uts, nil

}

