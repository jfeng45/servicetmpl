package dataservicefactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/datastorefactory"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/jfeng45/servicetmpl/dataservice/coursedata/couchdb"
	"github.com/jfeng45/servicetmpl/dataservice/coursedata/sqldb"
	"github.com/jfeng45/servicetmpl/tools/gdbc"
	"github.com/pkg/errors"
)

var courseDataServiceMap = map[string]dataservice.CourseDataInterface {
	datastorefactory.COUCHDB: &couchdb.CourseDataCouchdb{},
	datastorefactory.SQLDB:   &sqldb.CourseDataSql{},
}

// courseDataServiceFactory is an empty receiver for Build method
type courseDataServiceFactory struct {}

// GetCourseDataServiceInterface is an accessor for factoryBuilderMap
func GetCourseDataServiceInterface(key string) dataservice.CourseDataInterface {
	return courseDataServiceMap[key]
}

func (tdsf *courseDataServiceFactory) Build(c container.Container, dataConfig *configs.DataConfig) (DataServiceInterface, error) {
	logger.Log.Debug("courseDataServiceFactory")
	key := dataConfig.Code
	if COURSE_DATA != key {
		errMsg := COURSE_DATA + " in courseDataServiceFactory doesn't match key = " + key
		return nil, errors.New(errMsg)
	}

	dsc := dataConfig.DataStoreConfig
	dsi, err := datastorefactory.GetDataStoreFb(dsc.Code).Build(c, &dsc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	gdbc := dsi.(gdbc.Gdbc)
	gdi := GetCourseDataServiceInterface(dsc.Code)
	gdi.SetDB(gdbc)

	//c.Put(key, gdi)
	return gdi, nil
}
