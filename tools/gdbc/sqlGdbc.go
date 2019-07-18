// Package handles low level database access including transaction through *sql.Tx or *sql.DB

package gdbc

import (
	"context"
	"database/sql"
	"github.com/go-kivik/kivik"
)

// SqlGdbc (SQL Go database connection) is a wrapper for SQL database handler ( can be *sql.DB or *sql.Tx)
// It should be able to work with all SQL data that follows SQL standard.
type SqlGdbc interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	// If need transaction support, add this interface
	Transactioner
}

// SqlDBTx is the concrete implementation of sqlGdbc by using *sql.DB
type SqlDBTx struct {
	DB *sql.DB
}

// SqlConnTx is the concrete implementation of sqlGdbc by using *sql.Tx
type SqlConnTx struct {
	DB *sql.Tx
}

func (sdt *SqlDBTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return sdt.DB.Exec(query, args...)
}

func (sdt *SqlDBTx) Prepare(query string) (*sql.Stmt, error) {
	return sdt.DB.Prepare(query)
}

func (sdt *SqlDBTx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return sdt.DB.Query(query, args...)
}

func (sdt *SqlDBTx) QueryRow(query string, args ...interface{}) *sql.Row {
	return sdt.DB.QueryRow(query, args...)
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

// The followings are dummy implementation for noSqlGdbc.
// After implementing the following, SqlDBTx will also implements NoSqlGdbc interface.
// This making the courseDataServiceFactory possible.
// The functions will never be called, it just makes the interface available
// If you don't need something similar to courseDataServiceFactory, you can remove the following code.
func(sdt *SqlDBTx) QueryNoSql(ctx context.Context, ddoc string, view string)  (*kivik.Rows, error){
	return nil, nil
}

func(sdt *SqlDBTx) Put(ctx context.Context, docID string, doc interface{}, options ...kivik.Options) (rev string, err error) {
	return "", nil
}

func (sdt *SqlDBTx) Get(ctx context.Context, docID string, options ...kivik.Options) (*kivik.Row, error) {
	return nil, nil
}

func (sdt *SqlDBTx) Find(ctx context.Context, query interface{}) (*kivik.Rows, error) {
	return nil, nil
}

func(sdt *SqlDBTx) AllDocs(ctx context.Context, options ...kivik.Options ) (*kivik.Rows, error) {
	return nil, nil
}




