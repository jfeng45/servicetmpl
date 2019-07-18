package gdbc

import (
	"context"
	"database/sql"
	"github.com/go-kivik/kivik"
)

// NoSqlGdbc (NoSQL Go database connection) is a wrapper for NoSql database handler.
// Currently, it is just a POC instead of a mature implementation.
// To make it real, you need to remove kivik related types and need to do conversion between generic types and
// individual NoSql types
// It also doesn't includes all methods, just the one I need. You can add more when there is a need.
type NoSqlGdbc interface {
	// The method name of underline database was Query(), but since it conflicts with the name with Query() in SqlGdbc,
	// so have to change to a different name
	QueryNoSql(ctx context.Context,ddoc string, view string) (*kivik.Rows, error)
	Put(ctx context.Context, docID string, doc interface{}, options ...kivik.Options) (rev string, err error)
	Get(ctx context.Context, docID string, options ...kivik.Options) (*kivik.Row, error)
	Find(ctx context.Context, query interface{}) (*kivik.Rows, error)
	AllDocs(ctx context.Context, options ...kivik.Options ) (*kivik.Rows, error)
}

// NoSqlDB is the concrete implementation of NoSqlGdbc
type NoSqlDB struct {
	DB *kivik.DB
}

func (nsdb *NoSqlDB )QueryNoSql(ctx context.Context, ddoc string, view string)  (*kivik.Rows, error){
	// Query is the real name of the underline database method
	return nsdb.DB.Query(ctx ,ddoc , view)
}

func (nsdb *NoSqlDB )Put(ctx context.Context, docID string, doc interface{}, options ...kivik.Options) (rev string, err error) {
	return nsdb.DB.Put(ctx, docID, doc, options ...)
}

func (nsdb *NoSqlDB )Get(ctx context.Context, docID string, options ...kivik.Options) (*kivik.Row, error) {
	return nsdb.DB.Get(ctx, docID, options ...)
}

func (nsdb *NoSqlDB )Find(ctx context.Context, query interface{}) (*kivik.Rows, error) {
	return nsdb.DB.Find(ctx, query)
}

func (nsdb *NoSqlDB )AllDocs(ctx context.Context, options ...kivik.Options ) (*kivik.Rows, error) {
	return nsdb.DB.AllDocs(ctx, options ... )
}
// the followings are dummy implementation for SqlGdbc
// After implementing the following, NoSqlDB will also implements SqlGdbc interface.
// This making the courseDataServieFactory possible.
// The functions will never be called, it just makes the interface available
// If you don't need something similar to courseDataServiceFactory, you can remove the following code.
func (nsdb *NoSqlDB ) Exec(query string, args ...interface{}) (sql.Result, error) {
	return nil, nil
}

func (nsdb *NoSqlDB ) Prepare(query string) (*sql.Stmt, error) {
	return  nil, nil
}

func (nsdb *NoSqlDB ) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return  nil, nil
}

func (nsdb *NoSqlDB ) QueryRow(query string, args ...interface{}) *sql.Row {
	return  nil
}



