// Package mysql represents the CouchDB implementation of the user data persistence layer
package couchdb

import (
	"context"
	"github.com/go-kivik/kivik"
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/jfeng45/servicetmpl/model"
	"github.com/jfeng45/servicetmpl/tools/logger"
	"github.com/pkg/errors"
)

type UserDataCouchdb struct {
	DB *kivik.DB
}

// Create a view for "Find()", only need to run once.
// Should be created by Fauxton, but you may not know how to do it. To make it easy for you, put it in code.
// This function is created to make it easy to run the application, don't do it in production code
// When run for more than once, it will show error "Conflict: Document update conflict", just ignore it
func createView(dataStore *UserDataCouchdb) {
	rev, err := dataStore.DB.Put(context.TODO(), "_design/serviceConfigDesignDoc", map[string]interface{}{
		"_id": "_design/serviceConfigDesignDoc",
		"views": map[string]interface{}{
			"serviceConfigByID": map[string]interface{}{
				"map": "function(doc) {\n  if (doc.uid) {\n emit(doc.uid, doc);\n}\n}",
			},
		},
		"language":"javascript",
	})
	// For each rnu after first, it will throw an error because it already exist. Just ignore it.
	if err != nil {
		logger.Log.Errorf("err:%v\n" , err)
	}
	logger.Log.Debug("rev:", rev)
}

func (dataStore *UserDataCouchdb) Find(id int) (*model.User, error) {
	var err error
	// only need to it once
	createView(dataStore)
	rows, err := dataStore.DB.Query(context.TODO(), "_design/serviceConfigDesignDoc", "_view/serviceConfigByID", map[string]interface{}{"reduce": false},
		kivik.Options{"key":id})

	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	var user *model.User
	if rows.Next() {
		user = &model.User{}
		if err := rows.ScanValue(user); err != nil {
			return nil, errors.Wrap(err, "")
		}
	}
	logger.Log.Debugf("view:%+v", user)

	if rows.Err() != nil {
		return nil, errors.Wrap(rows.Err(), "")
	}
	return user, nil
}

//The simple version (no need for view) of Find() to get it work, it is kind cheating because it didn't use the parameter id.
//func (dataStore *UserDataCouchdb) Find(id int) (*model.User, error) {
//	_id :="80a9134c7dfa53f67f6be214e1000fa7"
//	row, err :=dataStore.DB.Get(context.TODO(), _id)
//	if err != nil {
//		return nil, errors.Wrap(err, "")
//	}
//	var user model.User
//	if err=row.ScanDoc(&user); err!=nil {
//		panic(err)
//	}
//	logger.Log.Debugf("user:", user)
//	return &user, nil
//}

func (dataStore *UserDataCouchdb) Remove(username string) (int64, error) {

	return 0, nil
}
func (dataStore *UserDataCouchdb) Update(user *model.User) (int64, error) {
	return 0, nil
}

func (dataStore *UserDataCouchdb) Insert(user *model.User) (*model.User, error) {
	return nil, nil
}
func (dataStore *UserDataCouchdb) FindAll() ([]model.User, error) {
	return []model.User{}, nil
}
func (dataStore *UserDataCouchdb) FindByName(name string) (*model.User, error) {
	return nil, nil
}

//CouchDB doesn't support transaction, do nothing
func (dataStore *UserDataCouchdb)TxEnd( txFunc func() error) error {
	return nil
}
//CouchDB doesn't support transaction, do nothing
func (dataStore *UserDataCouchdb)TxBegin() (dataservice.UserDataInterface, error) {
	return nil, nil
}
