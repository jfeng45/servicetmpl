package TxDataService

import (
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/jfeng45/servicetmpl/tools/gdbc"
)

// UserDataMySql is the MySQL implementation of UserDatainterface
type TxDataMySql struct {
	DB gdbc.Gdbc
}

func (dataStore *TxDataMySql)TxEnd( txFunc func() error) error {
	return dataStore.DB.TxEnd(txFunc)
}

func (dataStore *TxDataMySql)TxBegin() (dataservice.TxDataInterface, error) {

	gdbc, error :=dataStore.DB.TxBegin()
	dbts := TxDataMySql{gdbc}
	return &dbts, error
}
