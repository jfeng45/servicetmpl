package dataservicefactory

import (
	"github.com/go-kivik/kivik"
	//"github.com/flimzy/kivik"
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/datastorefactory"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/jfeng45/servicetmpl/dataservice/userdata/sql"
	"github.com/jfeng45/servicetmpl/dataservice/userdata/couchdb"
	"github.com/jfeng45/servicetmpl/tools/gdbc"
	"github.com/pkg/errors"
)
// userDataServiceFactory is a empty receiver for Build method
type userDataServiceFactory struct {}

func (udmf *userDataServiceFactory) Build(c container.Container, dataConfig *configs.DataConfig) (DataServiceInterface, error) {
	logger.Log.Debug("userDataServiceFactory")
	key := dataConfig.Code
	if USER_DATA != key {
		errMsg := USER_DATA + " in userDataServiceFactory doesn't match key = " + key
		return nil, errors.New(errMsg)
	}
	//if it is already in container, return
	if value, found := c.Get(key); found {
		return value.(dataservice.UserDataInterface), nil
	}

	dsc := dataConfig.DataStoreConfig
	dsi, err := datastorefactory.GetDataStoreFb(dsc.Code).Build(c, &dsc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	if dataConfig.DataStoreConfig.Code == datastorefactory.SQL {
		ds := dsi.(gdbc.SqlGdbc)
		udm := sql.UserDataSql{DB: ds}
		logger.Log.Debug("udm:", udm.DB)
		c.Put(key, &udm)
		return &udm, nil
	} else if dataConfig.DataStoreConfig.Code == datastorefactory.COUCHDB {
		ds := dsi.(*kivik.DB)
		udm := couchdb.UserDataCouchdb{DB: ds}
		c.Put(key, &udm)
		return &udm, nil
	}
	return nil, nil
}
