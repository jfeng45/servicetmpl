package userdataservicefactory

import (
	"github.com/go-kivik/kivik"
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/datastorefactory"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/jfeng45/servicetmpl/dataservice/userdata/couchdb"
	"github.com/pkg/errors"
)
// couchdbUserDataServiceFactory is a empty receiver for Build method
type couchdbUserDataServiceFactory struct {}

func (cudsf *couchdbUserDataServiceFactory) Build(c container.Container, dataConfig *configs.DataConfig) (dataservice.UserDataInterface, error) {
	logger.Log.Debug("couchdbUserDataServiceFactory")
	dsc := dataConfig.DataStoreConfig
	dsi, err := datastorefactory.GetDataStoreFb(dsc.Code).Build(c, &dsc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	ds := dsi.(*kivik.DB)
	udc := couchdb.UserDataCouchdb{DB: ds}
	return &udc, nil

}
