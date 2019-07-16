// Package handles low level database access including transaction through *sql.Tx or *sql.DB

package gdbc

import (
	"database/sql"
)

// Gdbc (Go database connection) is a wrapper for database handler ( can be *sql.DB or *sql.Tx) in order to handle both
type Gdbc interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	//if need transaction support, need to add this interface
	Transactioner
}

// DBTx is the concrete implementation of GDBC by using *sql.DB
type DBTx struct {
	DB *sql.DB
}

// ConnTx is the concrete implementation of GDBC by using *sql.Tx
type ConnTx struct {
	DB *sql.Tx
}

func (db *DBTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.DB.Exec(query, args...)
}

func (db *DBTx) Prepare(query string) (*sql.Stmt, error) {
	return db.DB.Prepare(query)
}

func (db *DBTx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.DB.Query(query, args...)
}

func (db *DBTx) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.DB.QueryRow(query, args...)
}

func (db *ConnTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.DB.Exec(query, args...)
}

func (db *ConnTx) Prepare(query string) (*sql.Stmt, error) {
	return db.DB.Prepare(query)
}

func (db *ConnTx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.DB.Query(query, args...)
}

func (db *ConnTx) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.DB.QueryRow(query, args...)
}




