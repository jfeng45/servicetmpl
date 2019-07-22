package datastorefactory

import (
	"context"
	couchdbKivid "github.com/go-kivik/couchdb"
	"github.com/go-kivik/kivik"
	//"github.com/flimzy/kivik"
	"github.com/jfeng45/servicetmpl/config"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/pkg/errors"
)

// couchdbFactory is receiver for Build method
type couchdbFactory struct{}

// implement Build method for CouchDB database
func (cf *couchdbFactory) Build(c container.Container, dsc *config.DataStoreConfig) (DataStoreInterface, error) {
	logger.Log.Debug("couchdbFactory")
	key := dsc.Code

	//if it is already in container, return
	if value, found := c.Get(key); found {
		logger.Log.Debug("found couchdb in container for key:", key)
		return value.(*kivik.DB), nil
	}
	// Don't know why needs adding the following line, because the driver is already registered in init() in couchdbKiv
	// however, not adding this, I got the error "unknown driver "couch" (forgotten import?)"
	kivik.Register(config.COUCHDB, &couchdbKivid.Couch{})

	client, err := kivik.New(context.TODO(), dsc.Code, dsc.UrlAddress)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	db, err := client.DB(context.TODO(), dsc.DbName)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	c.Put(key, db)
	return db, nil

}
