// Package handles low level database access including transaction through *sql.Tx or *sql.DB

package gdbc

import (
	"database/sql"
)

// SqlGdbc (Go database connection) is a wrapper for database handler ( can be *sql.DB or *sql.Tx) in order to handle both
// It should be able to work with all SQL data that follows SQL standard
type SqlGdbc interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	//if need transaction support, need to add this interface
	Transactioner
}

// SqlDBTx is the concrete implementation of GDBC by using *sql.DB
type SqlDBTx struct {
	DB *sql.DB
}

// SqlConnTx is the concrete implementation of GDBC by using *sql.Tx
type SqlConnTx struct {
	DB *sql.Tx
}

func (sdb *SqlDBTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return sdb.DB.Exec(query, args...)
}

func (sdb *SqlDBTx) Prepare(query string) (*sql.Stmt, error) {
	return sdb.DB.Prepare(query)
}

func (sdb *SqlDBTx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return sdb.DB.Query(query, args...)
}

func (sdb *SqlDBTx) QueryRow(query string, args ...interface{}) *sql.Row {
	return sdb.DB.QueryRow(query, args...)
}

func (sdb *SqlConnTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return sdb.DB.Exec(query, args...)
}

func (sdb *SqlConnTx) Prepare(query string) (*sql.Stmt, error) {
	return sdb.DB.Prepare(query)
}

func (sdb *SqlConnTx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return sdb.DB.Query(query, args...)
}

func (sdb *SqlConnTx) QueryRow(query string, args ...interface{}) *sql.Row {
	return sdb.DB.QueryRow(query, args...)
}




