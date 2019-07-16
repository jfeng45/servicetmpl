package dataservicefactory

import (
	"github.com/go-kivik/kivik"
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/datastorefactory"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/jfeng45/servicetmpl/dataservice/userdata/mysql"
	"github.com/jfeng45/servicetmpl/dataservice/userdata/couchdb"
	"github.com/jfeng45/servicetmpl/tools/gdbc"
	"github.com/pkg/errors"
)

type userDataServiceFactory struct {}

func (udmf *userDataServiceFactory) Build(c container.Container, userDataConfig *configs.UserDataConfig) (DataServiceInterface, error) {
	logger.Log.Debug("userDataServiceFactory")
	key := userDataConfig.Code
	if USER_DATA != key {
		errMsg := USER_DATA + " in userDataServiceFactory doesn't match key = " + key
		return nil, errors.New(errMsg)
	}
	//if it is already in container, return
	if value, found := c.Get(key); found {
		return value.(dataservice.UserDataInterface), nil
	}
	dsc := userDataConfig.DataStoreConfig
	dsi, err := datastorefactory.GetDataStoreFb(dsc.Code).Build(c, &dsc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	if userDataConfig.DataStoreConfig.Code == datastorefactory.MYSQL {
		ds := dsi.(gdbc.Gdbc)
		udm := mysql.UserDataMySql{DB: ds}
		logger.Log.Debug("udm:", udm.DB)
		c.Put(key, &udm)
		return &udm, nil
	} else if userDataConfig.DataStoreConfig.Code == datastorefactory.COUCHDB {
		ds := dsi.(kivik.DB)
		udm := couchdb.UserDataCouchdb{DB: &ds}
		c.Put(key, &udm)
		return &udm, nil
	}
	return nil, nil
}
