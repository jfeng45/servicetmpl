package gdbc

import (
	"fmt"
	"github.com/jfeng45/servicetmpl/tools"
)

// Transactioner is the transaction interface for database handler
type Transactioner interface {
	// Rollback a tranaction
	Rollback() error
	// Commit a transaction
	Commit() error
	// TxEnd commits a transaction if no errors, otherwise rollback
	// txFunc is the operations wrapped in a transaction
	TxEnd( txFunc func() error) error
	// TxBegin gets *sql.DB from receiver and return a Gdbc, which has a *sql.Tx
	TxBegin() (Gdbc, error)
}

// DB doesnt rollback, do nothing here
func (db *DBTx) Rollback() error {
	return nil
}

//DB doesnt commit, do nothing here
func (db *DBTx) Commit() error {
	return nil
}
// TransactionBegin starts a transaction
func (db *DBTx)TxBegin( ) (Gdbc, error) {
	fmt.Println("transaction begin")
	tx, err := db.DB.Begin()
	ct := ConnTx{tx}
	return &ct, err
}
// DB doesnt rollback, do nothing here
func (db *DBTx)TxEnd( txFunc func() error) error {
	return nil
}

func (db *ConnTx) TxEnd( txFunc func() error) error {
	var err error
	tx := db.DB

	defer func() {
		if p := recover(); p != nil {
			tools.Log.Debug("found p and rollback:", p)
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tools.Log.Debugf("found error and rollback:", err)
			tx.Rollback() // err is non-nil; don't change it
		} else {
			tools.Log.Debug("commit:")
			err = tx.Commit() // if Commit returns error update err with commit err
		}
	}()
	err = txFunc()
	return err
}
//*sql.Tx can't begin a transaction, transaction always begins with a *sql.DB
func (db *ConnTx) TxBegin( ) (Gdbc, error) {
	return nil, nil
}

func (db *ConnTx) Rollback() error {
	return db.DB.Rollback()
}

func (db *ConnTx) Commit() error {
	return db.DB.Commit()
}




