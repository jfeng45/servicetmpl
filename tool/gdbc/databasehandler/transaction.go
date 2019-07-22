package databasehandler

import (
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/tool/gdbc"
)

// DB doesn't rollback, do nothing here
func (cdt *SqlDBTx) Rollback() error {
	return nil
}

//DB doesnt commit, do nothing here
func (cdt *SqlDBTx) Commit() error {
	return nil
}

// TransactionBegin starts a transaction
func (cdt *SqlDBTx) TxBegin() (gdbc.SqlGdbc, error) {
	logger.Log.Debug("transaction begin")
	tx, err := cdt.DB.Begin()
	ct := SqlConnTx{tx}
	return &ct, err
}

// DB doesnt rollback, do nothing here
func (cdt *SqlDBTx) TxEnd(txFunc func() error) error {
	return nil
}

func (sct *SqlConnTx) TxEnd(txFunc func() error) error {
	var err error
	tx := sct.DB

	defer func() {
		if p := recover(); p != nil {
			logger.Log.Debug("found p and rollback:", p)
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			logger.Log.Debugf("found error and rollback:", err)
			tx.Rollback() // err is non-nil; don't change it
		} else {
			logger.Log.Debug("commit:")
			err = tx.Commit() // if Commit returns error update err with commit err
		}
	}()
	err = txFunc()
	return err
}

//*sql.Tx can't begin a transaction, transaction always begins with a *sql.DB
func (sct *SqlConnTx) TxBegin() (gdbc.SqlGdbc, error) {
	return nil, nil
}

func (sct *SqlConnTx) Rollback() error {
	return sct.DB.Rollback()
}

func (sct *SqlConnTx) Commit() error {
	return sct.DB.Commit()
}
