package dataservicefactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/datastorefactory"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/jfeng45/servicetmpl/dataservice/txdataservice"
	"github.com/jfeng45/servicetmpl/tools/gdbc"
	"github.com/pkg/errors"
)
// userDataServiceFactory is a empty receiver for Build method
type txDataServiceFactory struct {}

func (txdsf *txDataServiceFactory) Build(c container.Container, dataConfig *configs.DataConfig) (DataServiceInterface, error) {
	logger.Log.Debug("txDataServiceFactory")
	key := dataConfig.Code
	if TX_DATA != key {
		errMsg := TX_DATA + " in txDataServiceFactory doesn't match key = " + key
		return nil, errors.New(errMsg)
	}
	//if it is already in container, return
	if value, found := c.Get(key); found {
		return value.(dataservice.TxDataInterface), nil
	}

	dsc := dataConfig.DataStoreConfig
	dsi, err := datastorefactory.GetDataStoreFb(dsc.Code).Build(c, &dsc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	if dataConfig.DataStoreConfig.Code == datastorefactory.MYSQL {
		ds := dsi.(gdbc.Gdbc)
		tdm := txdataservice.TxDataMySql{ds}
		logger.Log.Debug("udm:", tdm.DB)
		c.Put(key, &tdm)
		return &tdm, nil
	} else {
		errMsg := "data store code " + dataConfig.DataStoreConfig.Code + " is not supported for transaction"
		return nil, errors.New(errMsg)
	}

}
