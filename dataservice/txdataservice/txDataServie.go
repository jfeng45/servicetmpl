// Package txdataservice represents transaction support on data service layer
package txdataservice

import (
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/jfeng45/servicetmpl/tools/gdbc"
)

// TxDataSql is the generic implementation for transaction for SQL database
// You only need to do it once for each SQL database
type TxDataSql struct {
	DB gdbc.SqlGdbc
}

func (tds *TxDataSql)TxEnd( txFunc func() error) error {
	return tds.DB.TxEnd(txFunc)
}

func (tds *TxDataSql)TxBegin() (dataservice.TxDataInterface, error) {

	sqlTx, error :=tds.DB.TxBegin()
	tdi:= TxDataSql{sqlTx}
	return &tdi, error
}
